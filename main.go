package main

import (
  "fmt"
  "net/http"

  "upperfile.com/database"
  "upperfile.com/routes"
)

func main() {
  database.Connect()
  routes.LoadRoutes()

  fmt.Println("#####################################")
  fmt.Println("#                                   #")
  fmt.Println("#  Server is starting on port 3100  #")
  fmt.Println("#                                   #")
  fmt.Println("#####################################")

  err := http.ListenAndServe(":3100", nil)
  if err != nil {
    panic(err)
  }
}
