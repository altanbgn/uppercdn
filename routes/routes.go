package routes

import (
  "log"
	"encoding/json"
	"net/http"
)

type Response struct {
  Status int `json:"status"`
  Message string `json:"message"`
}

func LoadRoutes() {
  http.Handle("/app", AppRouter())
  http.Handle("/api-key", APIKeyRouter())

  http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    responseObj := Response {
      Status: http.StatusOK,
      Message: "Health is OK!",
    }
    response, err := json.Marshal(responseObj)
    if err != nil {
      panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    num, err := w.Write(response)

    if err != nil {
      panic(err)
    }

    log.Println(num)
  })
}
