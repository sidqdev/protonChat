package storage

import "github.com/gorilla/sessions"

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key      = []byte("super-secret-key")
	Store    = sessions.NewCookieStore(key)
	Users    = UserStorage{}
	Messages = MessageStorage{}
	Updates  = UpdateStorage{}
)
