package storage

type Message struct {
	Text     string `json:"text"`     // only text
	FromUser string `json:"fromUser"` // means username
	ToUser   string `json:"toUser"`   // also means username
}

func (m *Message) IsBelongs(username1, username2 string) bool {
	return (m.FromUser == username1 && m.ToUser == username2) || (m.FromUser == username2 && m.ToUser == username1)
}

type MessageStorage struct {
	Messages []Message `json:"messages"`
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

func (u *UpdateStorage) SendMessage(message Message) {
	u.Messages = append(u.Messages, message)
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
	return updates
}
