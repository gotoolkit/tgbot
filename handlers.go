package tgbot

const (
	OnText = "/text"
)

type HandlerFunc func(*Message)

func (b *Bot) Handle(endpoint string, handler HandlerFunc) {
	b.handlers[endpoint] = handler
}

func (b *Bot) handle(endpoint string, m *Message) bool {
	handler, ok := b.handlers[endpoint]
	if !ok {
		return false
	}
	go handler(m)
	return true
}
