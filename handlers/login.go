package handlers

import (
	"encoding/json"
	"main/storage"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	session, _ := storage.Store.Get(r, "session")
	decoder := json.NewDecoder(r.Body)
	var u storage.User
	err := decoder.Decode(&u)
	if err != nil {
		http.Error(w, http.ErrBodyNotAllowed.Error(), http.StatusConflict)
	}
	status, userId := storage.Users.LoginUser(u)
	if status {
		session.Values["userId"] = userId
		http.Error(w, "accepted", http.StatusAccepted)
	} else {
		session.Values["userId"] = ""
		http.Error(w, "forbiden", http.StatusForbidden)
	}

	session.Save(r, w)
}
