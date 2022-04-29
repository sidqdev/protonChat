package main

import (
	"main/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/login", handlers.Login).Methods("POST")

	http.ListenAndServe("0.0.0.0:8060", route)
}
