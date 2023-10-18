package user

import (
  "github.com/gorilla/mux"
)

func LoadRoutes(r *mux.Router) {
  router := r.PathPrefix("/auth").Subrouter()

  router.
    Path("/login").
    Methods("POST").
    HandlerFunc(HandleLogin)

  router.
    Path("/register").
    Methods("POST").
    HandlerFunc(HandleRegister)
}
