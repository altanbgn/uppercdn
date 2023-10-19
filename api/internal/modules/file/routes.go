package file

import (
  // "upperfile.com/api/internal/middlewares"

	"github.com/gorilla/mux"
)

func LoadRoutes(r *mux.Router) {
	router := r.PathPrefix("/file").Subrouter()

	router.
    Path("/upload").
		Methods("POST").
		HandlerFunc(HandleUploadFile)
}
