package routes

import (
  "upperfile.com/handlers"

  "github.com/gorilla/mux"
)

func AppRouter() *mux.Router {
  r := mux.NewRouter()

  r.HandleFunc("/", handlers.HandleAppCreate).Methods("POST")
  r.HandleFunc("/{appId}", handlers.HandleAppGet).Methods("GET")
  r.HandleFunc("/{appId}", handlers.HandleAppUpdate).Methods("PUT")
  r.HandleFunc("/{appId}", handlers.HandleAppDelete).Methods("DELETE")

  return r
}
