package tgbot

import (
	"context"
	"errors"
	"log"
	"net/http"
)

type Bot struct {
	context context.Context
	stop    context.CancelFunc
	Token   string
	client  *http.Client
	Me      *User
	Updates chan Update
	Hook    Hook
}

// New build a bot with token
func New(token string, opts ...OptionFunc) (*Bot, error) {
	ctx, cancel := context.WithCancel(context.Background())
	bot := &Bot{
		context: ctx,
		stop:    cancel,
		Token:   token,
		Updates: make(chan Update, 50),
	}

	for _, opt := range opts {
		err := opt(bot)
		if err != nil {
			return nil, err
		}
	}

	if bot.client == nil {
		bot.client = http.DefaultClient
	}
	if bot.Hook == nil {
		bot.Hook = DefaultHook
	}
	me, err := bot.getMe()
	if err != nil {
		return nil, err
	}
	bot.Me = me

	return bot, nil
}

func (b *Bot) Start() error {
	if b.Hook == nil {
		return errors.New("no hooks")
	}
	ctx, cancel := context.WithCancel(context.Background())
	go b.Hook.Update(ctx, b)
	for {
		select {
		case up := <-b.Updates:
			b.processUpdate(&up)
		case <-b.context.Done():
			log.Println("cancel")
			cancel()
		case <-ctx.Done():
			log.Println("done")
			return nil
		}
	}
}

func (b *Bot) Stop() {
	b.stop()
}

func (b *Bot) processUpdate(up *Update) {
	log.Println("update")
}
