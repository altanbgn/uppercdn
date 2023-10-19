package apikey

import (
	"encoding/json"
	"net/http"
	"time"

	"upperfile.com/internal/db"
	"upperfile.com/internal/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type apikey struct {
	Key       string    `json:"key"`
	ExpireAt  time.Time `json:"expire_at"`
	ProjectID uuid.UUID `json:"project_id"`
}

func HandleAPIKeyCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	payload := apikey{}
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "BAD_REQUEST",
			"message": err.Error(),
		})

		return
	}

	err = db.DB.Exec(
		"INSERT INTO api_keys (key, expire_at, project_id) VALUES (?, ?, ?)",
		payload.Key,
		payload.ExpireAt,
    payload.ProjectID,
	).Error

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    _ = json.NewEncoder(w).Encode(map[string]string{
      "status":  "INTERNAL_SERVER_ERROR",
      "message": err.Error(),
    })

    return
  }

  w.WriteHeader(http.StatusCreated)
  _ = json.NewEncoder(w).Encode(map[string]string{
    "status":  "CREATED",
    "message": "API Key created successfully",
  })
}

func HandleAPIKeyList(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  var apiKeys []db.APIKey

	query := r.URL.Query()

	page := query.Get("page")
	perPage := query.Get("limit")

	if page == "" {
		page = "1"
	}

	if perPage == "" {
		perPage = "10"
	}

  _ = json.NewEncoder(w).Encode(apiKeys)
}

