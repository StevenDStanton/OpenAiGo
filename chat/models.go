package chat

import "encoding/json"

type Function struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type ToolCall struct {
	ID       string   `json:"id"`
	Type     string   `json:"type"`
	Function Function `json:"function"`
}

type Message struct {
	Content   string     `json:"content"`
	Role      ChatRole   `json:"role"`
	Name      string     `json:"name,omitempty"`
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
}

//TODO: Implement the following optional fields.
// - frequency_penalty
// - logit_bias
// - logprobs
// - top_logprobs
// - max_tokens
// - n
// - presence_penalty
// - seed
// - service_tier
// - stop
// - stream
// - stream_options
// - temperature
// - top_p
// - tools
// - tool_choice
// - parallel_tool_calls
// - user

type ChatRequest struct {
	Messages []Message `json:"messages"`
	Model    ChatModel `json:"model"`
}

type ChatResponse struct {
	ID                string    `json:"id"`
	Object            string    `json:"object"`
	Created           int64     `json:"created"`
	Model             string    `json:"model"`
	SystemFingerprint string    `json:"system_fingerprint"`
	Choices           []Choice  `json:"choices"`
	Usage             Usage     `json:"usage"`
	Error             *APIError `json:"error,omitempty"`
}

type Choice struct {
	Index        int              `json:"index"`
	Message      Message          `json:"message"`
	Logprobs     *json.RawMessage `json:"logprobs,omitempty"`
	FinishReason string           `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Service struct {
	apiKey string
}

type APIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param"`
	Code    string `json:"code"`
}
