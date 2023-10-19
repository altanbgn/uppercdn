package user

import (
	"encoding/json"
	"net/http"

	"upperfile.com/api/internal/db"
	"upperfile.com/api/internal/utils"
)

type LoginBody struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

	payload := LoginBody{}
  err := json.NewDecoder(r.Body).Decode(&payload)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    _ = json.NewEncoder(w).Encode(map[string]string{
      "status": "BAD_REQUEST",
      "message": err.Error(),
    })

    return
  }

  foundUser := &db.User{}
  db.DB.Where("username = ?", payload.Username).First(&foundUser)

  match, err := utils.ComparePasswordAndHash(payload.Password, foundUser.Password)
  if err != nil || !match {
    w.WriteHeader(http.StatusBadRequest)
    _ = json.NewEncoder(w).Encode(map[string]string{
      "status": "BAD_REQUEST",
      "message": "Invalid username or password",
    })

    return
  }

  token, err := utils.NewAccessToken(foundUser.ID.String())
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    _ = json.NewEncoder(w).Encode(map[string]string{
      "status": "INTERNAL_SERVER_ERROR",
      "message": err.Error(),
    })
  }

  w.WriteHeader(http.StatusInternalServerError)
  _ = json.NewEncoder(w).Encode(map[string]string{
    "status": "OK",
    "data": token,
  })
}

type RegisterBody struct {
	FirstName string `json:"firstname" form:"firstname"`
	LastName  string `json:"lastname" form:"lastname"`
	Email     string `json:"email" form:"email"`
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  payload := new(RegisterBody)

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    _ = json.NewEncoder(w).Encode(map[string]string{
      "status": "BAD_REQUEST",
      "message": err.Error(),
    })

    return
	}

  hashedPassword, err := utils.GenerateFromPassword(payload.Password, &utils.ArgonParams{
		Memory:      64 * 1024,
		Iterations:  1,
		Parallelism: 4,
		SaltLength:  16,
		KeyLength:   32,
	})

	if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    _ = json.NewEncoder(w).Encode(map[string]string{
      "status": "BAD_REQUEST",
      "message": err.Error(),
    })

    return
	}

	response := db.DB.Create(&db.User{
		Firstname: payload.FirstName,
		Lastname:  payload.LastName,
		Email:     payload.Email,
		Username:  payload.Username,
		Password:  hashedPassword,
	})

	if response.Error != nil {
    w.WriteHeader(http.StatusBadRequest)
    _ = json.NewEncoder(w).Encode(map[string]string{
      "status": "BAD_REQUEST",
      "message": response.Error.Error(),
    })

    return
	}

  w.WriteHeader(http.StatusOK)
  _ = json.NewEncoder(w).Encode(map[string]string{
    "status": "CREATED",
    "message": "User created successfully",
  })
}
