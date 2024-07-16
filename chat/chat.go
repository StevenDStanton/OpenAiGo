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
	if responseType == JSON {
		message := Message{
			Role:    System,
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

	// Debug statement to log the response body
	fmt.Println("Response Body:", string(responseBody))

	var chatResponse ChatResponse
	err = json.Unmarshal(responseBody, &chatResponse)
	if err != nil {
		// Attempt to unmarshal as error response
		var apiError APIError
		if jsonErr := json.Unmarshal(responseBody, &apiError); jsonErr == nil {
			return ChatResponse{}, fmt.Errorf("OpenAI API request failed with error: %s", apiError.Message)
		}
		return ChatResponse{}, fmt.Errorf("Failed to unmarshal response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		if chatResponse.Error != nil {
			return ChatResponse{}, fmt.Errorf("OpenAI API request failed with status code: %d, error: %s", resp.StatusCode, chatResponse.Error.Message)
		}
		return ChatResponse{}, fmt.Errorf("OpenAI API request failed with status code: %d, response body: %s", resp.StatusCode, responseBody)
	}

	return chatResponse, nil
}
