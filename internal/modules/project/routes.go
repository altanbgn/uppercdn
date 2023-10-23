package project

import (
  "upperfile.com/internal/middlewares"

	"github.com/gorilla/mux"
)

func LoadRoutes(r *mux.Router) {
	router := r.PathPrefix("/project").Subrouter()

  router.Use(middlewares.MustLoggedIn)

	router.
		Methods("POST").
		HandlerFunc(HandleProjectCreate)

  router.
    Path("/list").
    Methods("GET").
    HandlerFunc(HandleProjectList)

	router.
		Path("/{id}").
		Methods("GET", "PUT", "DELETE").
    HandlerFunc(HandleProjectID)
}
