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
	route.HandleFunc("/getMe", handlers.GetMe).Methods("GET")
	route.HandleFunc("/getChat", handlers.GetChat).Methods("GET")
	http.ListenAndServe("0.0.0.0:8060", route)
}
