package tgbot

// User object represents a Telegram user
type User struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

// ChatType represents one of the possible chat types.
type ChatType string

const (
	ChatPrivate    ChatType = "private"
	ChatGroup      ChatType = "group"
	ChatSuperGroup ChatType = "supergroup"
	ChatChannel    ChatType = "channel"
)

// Chat object represents a chat
type Chat struct {
	ID        int64    `json:"id"`
	Type      ChatType `json:"type"`
	Title     string   `json:"title"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Username  string   `json:"username"`
}
