package db

import (
  "log"
  "fmt"

  "upperfile.com/api/internal/config"

  "gorm.io/gorm"
  "gorm.io/gorm/logger"
  "gorm.io/driver/postgres"
)

var DB *gorm.DB

func Load() {
  var err error

  DB, err = gorm.Open(postgres.Open(config.Env.DB_URL), &gorm.Config{
    SkipDefaultTransaction: true,
    PrepareStmt: true,
    Logger: logger.Default.LogMode(logger.Silent),
  })
  if err != nil {
    log.Fatalln("Error connecting to database")
  }

  err = DB.AutoMigrate(
    &User{},
    &Project{},
  )
  if err != nil {
    log.Fatalln("Error migrating database")
  }

  fmt.Println("Loaded database")
}
