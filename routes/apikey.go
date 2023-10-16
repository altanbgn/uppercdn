package routes

import (
  "upperfile.com/handlers"

  "github.com/gorilla/mux"
)

func APIKeyRouter() *mux.Router {
  r := mux.NewRouter()

  r.HandleFunc("/", handlers.HandleAPIKeyCreate).Methods("POST")
  r.HandleFunc("/{appId}", handlers.HandleAPIKeyCreate).Methods("POST")

  return r
}
