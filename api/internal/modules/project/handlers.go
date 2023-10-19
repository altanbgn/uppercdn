package project

import (
	"encoding/json"
	"net/http"

	"upperfile.com/api/internal/db"
	"upperfile.com/api/internal/utils"

  "github.com/lib/pq"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
  AllowedUrls []string `json:"allowed_urls"`
}

func HandleProjectCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user_id := r.Context().Value(utils.UserContextKey)

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
		"INSERT INTO projects (name, description, allowed_urls, user_id) VALUES (?, ?, ?, ?)",
		payload.Name,
		payload.Description,
    pq.StringArray(payload.AllowedUrls),
		uuid.MustParse(user_id.(string)),
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

func HandleProjectList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user_id := r.Context().Value(utils.UserContextKey)
	query := r.URL.Query()

	page := query.Get("page")
	per_page := query.Get("perPage")

	if page == "" {
		page = "1"
	}

	if per_page == "" {
		per_page = "10"
	}

	foundProjects := []map[string]interface{}{}
	err := db.DB.
		Raw(
			"SELECT * FROM projects WHERE user_id = ? ORDER BY updated_at OFFSET (? - 1) * ? FETCH NEXT ? ROWS ONLY",
			user_id,
			page,
			per_page,
			per_page,
		).
		Scan(&foundProjects).
		Error

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "OK",
		"message": "Projects fetched successfully",
		"data":    foundProjects,
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

      foundProject["AllowedUrls"] = pq.StringArray(foundProject{"AllowedUrls"}).Value()

			w.WriteHeader(http.StatusCreated)
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  "OK",
				"message": "Project fetched successfully",
				"data": map[string]interface{}{
          AllowedUrls: pq.StringArray(foundProject.AllowedUrls).Value(),
        },
			})
		}

	case "PUT":
		{
			err = db.DB.
				Model(&db.Project{}).
				Updates(payload).
				Where("id = ?", id).
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
