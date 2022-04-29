package main

import (
	"log"
	"main/handlers"
	"main/storage"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	storage.Store.Options.Path = "/"
	storage.Store.Options.MaxAge = 0
	storage.Store.Options.Domain = ""

	log.Println("Init server")
	route := mux.NewRouter()
	route.HandleFunc("/login", handlers.Login).Methods("POST")
	route.HandleFunc("/logout", handlers.Logout).Methods("GET")
	route.HandleFunc("/getMe", handlers.GetMe).Methods("GET")
	route.HandleFunc("/getChat", handlers.GetChat).Methods("GET")
	http.ListenAndServe("localhost:8060", route)
}
