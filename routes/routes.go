package routes

import (
  "net/http"
)

func LoadRoutes() {
  http.Handle("/app", AppRouter())
  http.Handle("/api-key", APIKeyRouter())
}
