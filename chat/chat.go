package chat

type ChatModel string

type ChatRole string

type ResponseType string

const (
	GPT4o     ChatModel = "gpt-4o"
	GPT4Turbo ChatModel = "gpt-4-turbo"
	GPT4      ChatModel = "gpt-4"
	GPT35     ChatModel = "gpt-3.5-turbo"

	System    ChatRole = "system"
	User      ChatRole = "user"
	Assistant ChatRole = "assistant"
	Tool      ChatRole = "tool"

	Text ResponseType = "text"
	JSON ResponseType = "json_object"
)

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

func NewChatRequest(model ChatModel, messages []Message, responseType ResponseType) ChatRequest {
	if responseType == JSON {
		message := Message{
			Role:    System,
			Content: "Use Json to produce your response",
		}
		messages = append(messages, message)
	}
	return ChatRequest{
		Messages: messages,
		Model:    model,
	}
}
