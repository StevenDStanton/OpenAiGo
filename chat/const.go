package chat

// ChatModel represents the model used for chat
type ChatModel string

// ChatRole represents the role of a message sender in chat
type ChatRole string

// ResponseType represents the type of response expected from chat
type ResponseType string

// Models struct encapsulates chat models
type Models struct {
	GPT4o     ChatModel
	GPT4Turbo ChatModel
	GPT4      ChatModel
	GPT35     ChatModel
}

// Roles struct encapsulates chat roles
type Roles struct {
	System    ChatRole
	User      ChatRole
	Assistant ChatRole
	Tool      ChatRole
}

// Types struct encapsulates response types
type Types struct {
	Text ResponseType
	JSON ResponseType
}

// Initialize grouped constants
var (
	Model = Models{
		GPT4o:     "gpt-4o",
		GPT4Turbo: "gpt-4-turbo",
		GPT4:      "gpt-4",
		GPT35:     "gpt-3.5-turbo",
	}
	Role = Roles{
		System:    "system",
		User:      "user",
		Assistant: "assistant",
		Tool:      "tool",
	}
	Type = Types{
		Text: "text",
		JSON: "json_object",
	}
)
