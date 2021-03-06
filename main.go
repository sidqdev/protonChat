package main

import (
	"log"
	"main/handlers"
	"main/storage"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	storage.Users.Load()
	storage.Messages.Load()
	storage.Updates.Load()

	log.Println("Init server")
	route := mux.NewRouter()
	route.HandleFunc("/login", handlers.Login).Methods("POST")
	route.HandleFunc("/logout", handlers.Logout).Methods("GET")
	route.HandleFunc("/getMe", handlers.GetMe).Methods("GET")
	route.HandleFunc("/getChat", handlers.GetChat).Methods("GET")
	route.HandleFunc("/sendMessage", handlers.SendMessage).Methods("PUT")
	route.HandleFunc("/getUpdates", handlers.GetUpdates).Methods("GET")
	http.ListenAndServe("localhost:8060", route)
}
