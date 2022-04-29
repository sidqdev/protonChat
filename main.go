package main

import (
	"log"
	"main/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Init server")
	route := mux.NewRouter()
	route.HandleFunc("/login", handlers.Login).Methods("POST")
	http.ListenAndServe("0.0.0.0:8060", route)
}
