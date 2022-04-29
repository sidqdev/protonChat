package handlers

import (
	"main/storage"
	"net/http"
)

func GetChat(w http.ResponseWriter, r *http.Request) {
	session, _ := storage.Store.Get(r, "session")
	login := r.URL.Query().Get("login")
	_ = login
	session.Save(r, w)
}
