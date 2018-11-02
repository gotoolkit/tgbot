package tgbot

// User represents a Telegram user
type User struct {
	ID        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
}

// ChatType represents one of the possible chat types.
type ChatType string

const (
	ChatPrivate    ChatType = "private"
	ChatGroup      ChatType = "group"
	ChatSuperGroup ChatType = "supergroup"
	ChatChannel    ChatType = "channel"
)

// Chat represents a chat
type Chat struct {
	ID   int64    `json:"id"`
	Type ChatType `json:"type"`
}

// Update represents an incoming update.
type Update struct {
	ID      int      `json:"update_id"`
	Message *Message `json:"message,omitempty"`
}

// Message represents a message.
type Message struct {
	ID       int   `json:"message_id"`
	From     *User `json:"from"`
	Unixtime int64 `json:"date"`
	Chat     *Chat `json:"chat"`
}
