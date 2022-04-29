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
}
