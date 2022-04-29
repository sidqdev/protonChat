package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Message struct {
	Text     string `json:"text"`         // only text
	FromUser string `json:"fromUsername"` // means username
	ToUser   string `json:"toUsername"`   // also means username
}

func (m *Message) IsBelongs(username1, username2 string) bool {
	return (m.FromUser == username1 && m.ToUser == username2) || (m.FromUser == username2 && m.ToUser == username1)
}

type MessageStorage struct {
	Messages []Message `json:"messages"`
}

func (us *MessageStorage) Save() {
	content, err := json.Marshal(us)
	if err != nil {
		log.Println(err)
		return
	}
	err = ioutil.WriteFile("MessageStorage.json", content, 0644)
	if err != nil {
		log.Println(err)
	}
}

func (us *MessageStorage) Load() {
	content, err := ioutil.ReadFile("MessageStorage.json")
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(content, &us)
	if err != nil {
		log.Println(err)
	}
}

func (m *MessageStorage) GetMessages(fromUser, toUser string) []Message {
	messages := []Message{}
	for _, message := range m.Messages {
		if message.IsBelongs(fromUser, toUser) {
			messages = append(messages, message)
		}
	}
	return messages
}

func (m *MessageStorage) SendMessage(message Message) {
	m.Messages = append(m.Messages, message)
	Updates.SendMessage(message)
}

type UpdateStorage struct {
	Messages []Message `json:"messages"`
}

func (us *UpdateStorage) Save() {
	content, err := json.Marshal(us)
	if err != nil {
		log.Println(err)
		return
	}
	err = ioutil.WriteFile("UpdateStorage.json", content, 0644)
	if err != nil {
		log.Println(err)
	}
}

func (us *UpdateStorage) Load() {
	content, err := ioutil.ReadFile("UpdateStorage.json")
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(content, &us)
	if err != nil {
		log.Println(err)
	}
}

func (u *UpdateStorage) SendMessage(message Message) {
	u.Messages = append(u.Messages, message)
	u.Save()
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
	u.Save()
	return updates
}
