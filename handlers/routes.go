package handlers

import (
  //"net/http"

  "github.com/gorilla/mux"
)

func HandleRoutes(r *mux.Router){
  r.HandleFunc("/login", loginHandler)
  r.HandleFunc("/healthTest", healthTestHandler)
}
