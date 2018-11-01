package tgbot

import "net/http"

type Bot struct {
	Token  string
	client *http.Client
}

type OptionFunc func(*Bot) error

// NewBot build a bot with token
func NewBot(token string, opts ...OptionFunc) (*Bot, error) {

	bot := &Bot{Token: token}
	for _, opt := range opts {
		err := opt(bot)
		if err != nil {
			return nil, err
		}
	}

	if bot.client == nil {
		bot.client = http.DefaultClient
	}

	user, err := bot.getMe()

	return bot, nil
}
