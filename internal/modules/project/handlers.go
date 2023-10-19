package project

import (
	"encoding/json"
	"net/http"

	"upperfile.com/internal/db"
	"upperfile.com/internal/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func HandleProjectCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId := r.Context().Value(utils.UserContextKey)

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

	err = db.DB.Exec(
		"INSERT INTO projects (name, description, user_id) VALUES (?, ?, ?)",
		payload.Name,
		payload.Description,
		uuid.MustParse(userId.(string)),
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
	case "GET":
		{
			foundProject := map[string]interface{}{}
			err = db.DB.
				Raw("SELECT * FROM projects WHERE id = ?", id).
				Scan(&foundProject).
				Error

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_ = json.NewEncoder(w).Encode(map[string]string{
					"status":  "INTERNAL_SERVER_ERROR",
					"message": err.Error(),
				})

				return
			}

			w.WriteHeader(http.StatusCreated)
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  "OK",
				"message": "Project fetched successfully",
				"data":    foundProject,
			})
		}

	case "PUT":
		{
      err = db.DB.
        Model(&db.Project{}).
        Where("id = ?", id).
        Updates(payload).
        Error

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_ = json.NewEncoder(w).Encode(map[string]string{
					"status":  "INTERNAL_SERVER_ERROR",
					"message": err.Error(),
				})

				return
			}

			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(map[string]string{
				"status":  "OK",
				"message": "Project updated successfully",
			})
		}

	case "DELETE":
		{
			err = db.DB.
				Exec("DELETE FROM projects WHERE id = ?", id).
				Error

      if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        _ = json.NewEncoder(w).Encode(map[string]string{
          "status":  "INTERNAL_SERVER_ERROR",
          "message": err.Error(),
        })

        return
      }

      w.WriteHeader(http.StatusOK)
      _ = json.NewEncoder(w).Encode(map[string]string{
        "status":  "OK",
        "message": "Project deleted successfully",
      })
		}
	}
}
