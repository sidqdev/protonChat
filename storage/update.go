package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

type UpdateStorage struct {
	Messages []Message `json:"messages"`
	mu       sync.Mutex
}

func (u *UpdateStorage) Save() {
	u.mu.Lock()
	content, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
		return
	}
	err = ioutil.WriteFile("UpdateStorage.json", content, 0644)
	if err != nil {
		log.Println(err)
	}
	u.mu.Unlock()
}

func (u *UpdateStorage) Load() {
	content, err := ioutil.ReadFile("UpdateStorage.json")
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(content, &u)
	if err != nil {
		log.Println(err)
	}
}

func (u *UpdateStorage) SendMessage(message Message) {
	u.Messages = append(u.Messages, message)
	go u.Save()
}

func (u *UpdateStorage) GetUpdates(username string) []Message {
	updates := []Message{}
	i := 0
	for i < len(u.Messages) {
		if u.Messages[i].ToUser == username {
			updates = append(updates, u.Messages[i])
			u.Messages = append(u.Messages[:i], u.Messages[i+1:]...)
		} else {
			i += 1
		}
	}
	go u.Save()
	return updates
}
