package tgbot

import "net/http"

type OptionFunc func(*Bot) error

func WithClient(client *http.Client) OptionFunc {
	return func(b *Bot) error {
		b.client = client
		return nil
	}
}
