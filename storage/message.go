package storage

type Message struct {
	Text     string `json:"text"`     // only text
	FromUser string `json:"fromUser"` // means username
	ToUser   string `json:"toUser"`   // also means username
}

type MessageStorage struct {
	Messages []Message `json:"messages"`
}

func (m *MessageStorage) GetMessages(fromUser, toUser string) []Message {
	return nil
}
