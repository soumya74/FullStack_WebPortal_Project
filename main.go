package main

import (
  "net/http"
  "log"
  "fmt"

  "github.com/gorilla/mux"
  "github.com/FullStack_WebPortal_Project/handlers"
)

func main(){
  router := mux.NewRouter().StrictSlash(true)
  portalRouter := router.PathPrefix("/portal").Subrouter()
  handlers.HandleRoutes(portalRouter)
  fmt.Println("listening at localhost:10001")
  log.Fatal(http.ListenAndServe(":10001", router))
}
