package project

import (
	"encoding/json"
	"net/http"

	"upperfile.com/internal/db"
	"upperfile.com/internal/utils"

	"github.com/gorilla/mux"
)

type project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func HandleProjectCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  userId := r.Context().Value(utils.UserContextKey)

  foundUser := &db.User{}
  userResponse := db.DB.First(foundUser).Where("id = ?", userId)
  if userResponse.Error != nil {
    w.WriteHeader(http.StatusNotFound)
    _ = json.NewEncoder(w).Encode(map[string]string{
      "status":  "NOT_FOUND",
      "message": userResponse.Error.Error(),
    })

    return
  }

	payload := project{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "BAD_REQUEST",
			"message": err.Error(),
		})

		return
	}

	response := db.DB.Create(&db.Project{
		Name:        payload.Name,
		Description: payload.Description,
	})

	if response.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "INTERNAL_SERVER_ERROR",
			"message": response.Error.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status":  "CREATED",
		"message": "Project created successfully",
	})
}

func HandleProjectID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

  vars := mux.Vars(r)
  id := vars["id"]

	payload := project{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "BAD_REQUEST",
			"message": err.Error(),
		})

		return
	}

  switch r.Method {
    case "GET": {
      foundProject := &db.Project{}
      response := db.DB.First(foundProject).Where("id = ?", id)

      if response.Error != nil {
        w.WriteHeader(http.StatusBadRequest)
        _ = json.NewEncoder(w).Encode(map[string]string{
          "status":  "INTERNAL_SERVER_ERROR",
          "message": response.Error.Error(),
        })

        return
      }

      w.WriteHeader(http.StatusOK)
      _ = json.NewEncoder(w).Encode(map[string]interface{}{
        "status":  "OK",
        "message": "Project fetched successfully",
        "data": foundProject,
      })

    }
  }
}
