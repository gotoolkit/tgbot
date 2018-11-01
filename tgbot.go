package tgbot

import "net/http"

type Bot struct {
	Token  string
	client *http.Client
	Me     *User
}

// New build a bot with token
func New(token string, opts ...OptionFunc) (*Bot, error) {

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

	me, err := bot.getMe()
	if err != nil {
		return nil, err
	}
	bot.Me = me

	return bot, nil
}
