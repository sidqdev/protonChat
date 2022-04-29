package storage

type Message struct {
	Text     string `json:"text"`     // only text
	FromUser string `json:"FromUser"` // means login
	ToUser   string `json:"ToUser"`   // also means login
}
