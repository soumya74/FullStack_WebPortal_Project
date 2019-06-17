package handlers

import(
  "fmt"
  "net/http"

  //"github.com/gorilla/mux"
)

// loginHandler renders login page
func loginHandler(w http.ResponseWriter, r *http.Request){
  fmt.Println("login page hit")
}
