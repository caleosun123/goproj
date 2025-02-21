package main

import (
  "fmt"
  "net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
 /* if r.URL.Path != "/" {
    errorHandler(w, r, http.StatusNotFound)
    return
  }
*/
  http.ServeFile(w, r, "static/index.html")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
  w.WriteHeader(status)
  if status == http.StatusNotFound {
    fmt.Fprint(w, "404 Error: Not found")
  }
}

func main() {
  http.HandleFunc("/", indexHandler)
  fmt.Println("Starting server at port 8080...")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    fmt.Println("Error starting server:", err)
  }
}
