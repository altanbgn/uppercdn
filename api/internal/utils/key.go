package utils

import (
  "crypto/rand"
  "encoding/hex"
)

func GenerateApiKey(length int) (string, error) {
  key := make([]byte, length)
  _, err := rand.Read(key)
  if err != nil {
    return "", err
  }

  return hex.EncodeToString(key), nil
}
