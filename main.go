package main

import (
  "fmt"
  "net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, world!")
}

func main() {
  http.HandleFunc("/", indexHandler)
  fmt.Println("Starting server at port 8080...")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    fmt.Println("Error starting server:", err)
  }
}
