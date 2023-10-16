package handlers

import (
	"encoding/json"
	"fmt"
  "io"
	"net/http"

  "upperfile.com/database"
  "upperfile.com/models"
  "upperfile.com/utils"
)

type APIKeyResponse struct {
  status int
  message string
}

func HandleAPIKeyCreate(w http.ResponseWriter, r *http.Request) {
  body, err := io.ReadAll(r.Body)
  if err != nil {
    panic(err)
  }

  fmt.Println(body)

  apiKey, err := utils.GenerateApiKey(32)
  if err != nil {
    panic(err)
  }

  database.Instance.Create(&models.APIKey{
    Key: apiKey,
    ExpiryDate: "2021-01-01",
  })

  response := APIKeyResponse {
    status: http.StatusCreated,
    message: "API Key created",
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(response)
}
