package structs

import "main/storage"

type GetChatResponse struct {
	MyUsername string            `json:"myUsername"`
	Username   string            `json:"username"`
	Messages   []storage.Message `json:"messages"`
}
