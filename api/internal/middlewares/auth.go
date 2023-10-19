package middlewares

import (
  "context"
  "net/http"
	"strings"
  "encoding/json"

  "upperfile.com/api/internal/utils"
)

func VerifyAuth(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    header := r.Header.Get("Authorization")
    headerSplit := strings.Split(header, " ")

    if len(headerSplit) < 2 {
      next.ServeHTTP(w, r)
      return
    }

    token := headerSplit[1]
    tokenData := utils.ParseAccessToken(token)

    ctx := context.WithValue(r.Context(), utils.UserContextKey, tokenData.ID)
    r = r.WithContext(ctx)

    next.ServeHTTP(w, r)
  })
}

func MustLoggedIn(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    userContext := r.Context().Value(utils.UserContextKey)

    if userContext == nil {
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusUnauthorized)
      _ = json.NewEncoder(w).Encode(map[string]string{
        "status":  "UNAUTHORIZED",
        "message": "You must be logged in to access this resource",
      })

      return
    }

    next.ServeHTTP(w, r)
  })
}
