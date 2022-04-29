package handlers

import (
	"encoding/json"
	"log"
	"main/storage"
	"main/structs"
	"net/http"
)

func GetChat(w http.ResponseWriter, r *http.Request) {
	log.Println("Get chat")
	session, _ := storage.Store.Get(r, "user-storage")
	userIdInterface := session.Values["userId"]
	if userIdInterface == nil {
		http.Error(w, "please login before get chat", http.StatusForbidden)
		return
	}
	userId := userIdInterface.(string)
	myUsername, status := storage.Users.GetUserName(userId)
	if !status {
		http.Error(w, "please login before get chat", http.StatusForbidden)
		return
	}
	username := r.URL.Query().Get("username")
	if status := storage.Users.CheckUserExist(username); !status {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	messages := storage.Messages.GetMessages(myUsername, username)

	respone := structs.GetChatResponse{MyUsername: myUsername, Username: username, Messages: messages}
	json.NewEncoder(w).Encode(respone)
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("Send message")
	session, _ := storage.Store.Get(r, "user-storage")
	userIdInterface := session.Values["userId"]
	if userIdInterface == nil {
		http.Error(w, "please login before send message", http.StatusForbidden)
		return
	}
	userId := userIdInterface.(string)
	myUsername, status := storage.Users.GetUserName(userId)
	if !status {
		http.Error(w, "please login before send message", http.StatusForbidden)
		return
	}
	var message = storage.Message{}
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, http.ErrBodyNotAllowed.Error(), http.StatusConflict)
		return
	}
	message.FromUser = myUsername
	storage.Messages.SendMessage(message)
}

func GetUpdates(w http.ResponseWriter, r *http.Request) {
	log.Println("Get updates")
	session, _ := storage.Store.Get(r, "user-storage")
	userIdInterface := session.Values["userId"]
	if userIdInterface == nil {
		http.Error(w, "please login before get updates", http.StatusForbidden)
		return
	}
	userId := userIdInterface.(string)
	myUsername, status := storage.Users.GetUserName(userId)
	if !status {
		http.Error(w, "please login before get updates", http.StatusForbidden)
		return
	}
	updates := storage.Updates.GetUpdates(myUsername)
	json.NewEncoder(w).Encode(updates)
}
