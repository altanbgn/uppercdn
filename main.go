package main

import (
  "net/http"

  "upperfile.com/database"
  "upperfile.com/routes"
)

func main() {
  database.Connect()
  routes.LoadRoutes()

  err := http.ListenAndServe(":3100", nil)

  if err != nil {
    panic(err)
  }
}
