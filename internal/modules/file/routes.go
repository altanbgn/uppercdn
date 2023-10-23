package file

import (
  "net/http"

	"github.com/gorilla/mux"
)

func LoadRoutes(r *mux.Router) {
	router := r.PathPrefix("/file").Subrouter()

	router.
    Path("/upload").
		Methods("POST").
		HandlerFunc(HandleUploadFile)

  router.
    Handle(
      "/storage",
      http.StripPrefix(
        "/file/download",
        http.FileServer(http.Dir("./../../../../storage"),
      ),
    ))
}
