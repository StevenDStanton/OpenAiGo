package openaiclient

import "github.com/StevenDStanton/openaigo/chat"

type OpenAIClient struct {
	Chat *chat.Service
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
	return &OpenAIClient{
		Chat: chat.NewChatService(apiKey),
	}
}
