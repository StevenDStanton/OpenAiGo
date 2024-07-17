package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func NewChatService(apiKey string) *Service {
	return &Service{apiKey: apiKey}
}

func (s *Service) NewChatRequest(model ChatModel, messages []Message, responseType ResponseType) (ChatResponse, error) {
	if responseType == Type.JSON {
		message := Message{
			Role:    Role.System,
			Content: "Use Json to produce your response",
		}
		messages = append(messages, message)
	}
	chatRequest := ChatRequest{
		Messages: messages,
		Model:    model,
	}
	chatResponse, err := sendChatRequest(chatRequest, s)
	return chatResponse, err
}

func sendChatRequest(chatRequest ChatRequest, s *Service) (ChatResponse, error) {
	requestBody, err := json.Marshal(chatRequest)
	if err != nil {
		return ChatResponse{}, err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return ChatResponse{}, err
	}

	req.Header.Set("Authorization", "Bearer "+s.apiKey)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 90 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return ChatResponse{}, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ChatResponse{}, err
	}

	fmt.Println(resp.StatusCode)
	var chatResponse ChatResponse
	err = json.Unmarshal(responseBody, &chatResponse)
	if err != nil {
		return ChatResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return ChatResponse{}, fmt.Errorf("openAI API request failed with status code: %d, response body: %s", resp.StatusCode, chatResponse.Error.Message)
	}

	return chatResponse, nil

}
