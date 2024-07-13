package chat

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
