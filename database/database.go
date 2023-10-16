package database

import (
  "upperfile.com/models"

  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

var Instance *gorm.DB

func Connect() {
  db, err := gorm.Open(sqlite.Open("api-keys.db"), &gorm.Config{})
  if err != nil {
    panic("Failed to connect database")
  }

  err = db.AutoMigrate(&models.APIKey{})

  if err != nil {
    panic("Failed to migrate database")
  }

  Instance = db
}

