package handlers

import (
  "net/http"
)

func HandleAppCreate(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello World!"))
}
