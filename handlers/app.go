package handlers

import (
  "fmt"
  "io"
  "net/http"
)

func HandleAppCreate(w http.ResponseWriter, r *http.Request) {
  faq, err := io.WriteString(w, "Hello World!")
  if err != nil {
    panic(err)
  }

  fmt.Println(faq)
}

func HandleAppGet(w http.ResponseWriter, r *http.Request) {
  faq, err := io.WriteString(w, "Hello World!")
  if err != nil {
    panic(err)
  }

  fmt.Println(faq)
}

func HandleAppUpdate(w http.ResponseWriter, r *http.Request) {
  faq, err := io.WriteString(w, "Hello World!")
  if err != nil {
    panic(err)
  }

  fmt.Println(faq)
}

func HandleAppDelete(w http.ResponseWriter, r *http.Request) {
  faq, err := io.WriteString(w, "Hello World!")
  if err != nil {
    panic(err)
  }

  fmt.Println(faq)
}
