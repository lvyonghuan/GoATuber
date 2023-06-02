package helpers

import "github.com/google/uuid"

// OpenAIConversation is the struct that represents the json file
// the actual struct has many more fields, but we only need the ones below
type OpenAIConversation struct {
	Title   string                   `json:"title"`
	Mapping map[string]*Conversation `json:"mapping"`
}

type Conversation struct {
	Message *Message `json:"message"`
}

type Message struct {
	ID         uuid.UUID `json:"id"`
	Author     *Author   `json:"author"`
	CreateTime float64   `json:"create_time"`
	Content    *Content  `json:"content"`
}

type Author struct {
	Role string `json:"role"`
}

type Content struct {
	Parts []string `json:"parts"`
}
