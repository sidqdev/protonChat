package handlers

import (
	"encoding/json"
	"log"
	"main/storage"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	session, _ := storage.Store.Get(r, "user-storage")
	decoder := json.NewDecoder(r.Body)
	var u storage.User
	err := decoder.Decode(&u)
	if err != nil {
		http.Error(w, http.ErrBodyNotAllowed.Error(), http.StatusConflict)
		log.Println("Login error")
		return
	}
	log.Println("Login", u.UserName, u.Password)
	status, userId := storage.Users.LoginUser(u)
	if status {
		session.Values["userId"] = userId
		session.Save(r, w)
		http.Error(w, "accepted", http.StatusAccepted)
	} else {
		session.Values["userId"] = ""
		session.Save(r, w)
		http.Error(w, "forbiden", http.StatusForbidden)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	log.Println("Logout")
	session, _ := storage.Store.Get(r, "user-storage")
	userIdInterface := session.Values["userId"]
	if userIdInterface == nil {
		http.Error(w, "please login before logout", http.StatusForbidden)
		return
	}
	userId := userIdInterface.(string)
	storage.Users.Logout(userId)
	session.Values["userId"] = ""
	session.Save(r, w)
}

func GetMe(w http.ResponseWriter, r *http.Request) {
	log.Println("Get me")
	session, _ := storage.Store.Get(r, "user-storage")
	userIdInterface := session.Values["userId"]
	if userIdInterface == nil {
		http.Error(w, "please login before get chat", http.StatusForbidden)
		return
	}
	userId := userIdInterface.(string)
	username, status := storage.Users.GetUserName(userId)
	if status {
		http.Error(w, username, http.StatusOK)
	} else {
		http.Error(w, "forbiden", http.StatusForbidden)
	}
}
