package file

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
  "io"

  "upperfile.com/internal/config"
)

func HandleUploadFile(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  query := r.URL.Query()
  projectId := query.Get("projectId")

  if projectId == "" {
    w.WriteHeader(http.StatusBadRequest)
    _ = json.NewEncoder(w).Encode(map[string]string{
      "status": "BAD_REQUEST",
      "error": "Specify project id",
    })
  }

  file, fileHeader, err := r.FormFile("file")
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    _ = json.NewEncoder(w).Encode(map[string]string{
      "status": "BAD_REQUEST",
      "error": "Invalid file",
    })

    return
  }

  if fileHeader.Size > config.Env.MAX_FILE_SIZE {
    w.WriteHeader(http.StatusBadRequest)
    _ = json.NewEncoder(w).Encode(map[string]string{
      "status": "BAD_REQUEST",
      "error": "File size exceeded",
    })

    return
  }

  defer file.Close()

  err = os.MkdirAll(fmt.Sprintf("./../../../storage/%s", projectId), os.ModePerm)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    _ = json.NewEncoder(w).Encode(map[string]string{
      "status": "INTERNAL_SERVER_ERROR",
      "error": "Failed to create directory",
    })

    return
  }

  dst, err := os.Create(fmt.Sprintf("./../../../storage/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    _ = json.NewEncoder(w).Encode(map[string]string{
      "status": "INTERNAL_SERVER_ERROR",
      "error": "Failed to create file",
    })

    return
  }

  defer dst.Close()

  _, err = io.Copy(dst, file)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    _ = json.NewEncoder(w).Encode(map[string]string{
      "status": "INTERNAL_SERVER_ERROR",
      "error": "Failed to copy file",
    })

    return
  }

  w.WriteHeader(http.StatusOK)
  _ = json.NewEncoder(w).Encode(map[string]string{
    "status": "OK",
    "message": "File uploaded successfully",
  })
}
