package app

import (
	"fmt"
  "net/http"
  "encoding/json"

	"upperfile.com/internal/middlewares"
	"upperfile.com/internal/modules/user"
	"upperfile.com/internal/modules/project"

	"github.com/gorilla/mux"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func LoadRoutes(r *mux.Router) {
	v1 := r.PathPrefix("/v1").Subrouter()
	v1.Use(middlewares.VerifyAuth)

	user.LoadRoutes(v1)
  project.LoadRoutes(v1)

	v1.Path("/health").Methods("GET").HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    _ = json.NewEncoder(w).Encode(map[string]string{
      "status":  "200",
      "message": "OK",
    })
  })

	fmt.Println("Loaded routes")
}
