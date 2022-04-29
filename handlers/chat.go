package handlers

import (
	"encoding/json"
	"main/storage"
	"main/structs"
	"net/http"
)

func GetChat(w http.ResponseWriter, r *http.Request) {
	session, _ := storage.Store.Get(r, "session")
	username := r.URL.Query().Get("username")
	userId := session.Values["userId"].(string)
	myUsername, status := storage.Users.GetUserName(userId)
	if !status {
		http.Error(w, "please login before get chat", http.StatusForbidden)
		return
	}
	if status := storage.Users.CheckUserExist(username); !status {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	messages := storage.Messages.GetMessages(myUsername, username)

	respone := structs.GetChatResponse{MyUsername: myUsername, Username: username, Messages: messages}
	json.NewEncoder(w).Encode(respone)

	session.Save(r, w)
}
