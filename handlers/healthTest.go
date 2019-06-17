package handlers

import(
  "fmt"
  "net/http"
)

func healthTestHandler(w http.ResponseWriter, r *http.Request){
  fmt.Println("health test okay")
}
