package handlers

import (
	"log"
	"main/storage"
	"net/http"
)

func GetChat(w http.ResponseWriter, r *http.Request) {
	session, _ := storage.Store.Get(r, "session")
	username := r.URL.Query().Get("username")
	userId := session.Values["userId"].(string)
	myUserName, status := storage.Users.GetUserName(userId)
	if !status {
		http.Error(w, "please login before get chat", http.StatusForbidden)
		return
	}
	if status := storage.Users.CheckUserExist(username); !status {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	log.Println(myUserName)
	session.Save(r, w)
}
