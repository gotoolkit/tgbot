package tgbot

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Bot struct {
	ctx      context.Context
	stop     context.CancelFunc
	Token    string
	client   *http.Client
	Me       *User
	Hook     Hook
	Updates  chan Update
	Timeout  time.Duration
	handlers map[string]HandlerFunc
}

// New build a bot with token
func New(token string, opts ...OptionFunc) (*Bot, error) {
	ctx, cancel := context.WithCancel(context.TODO())
	bot := &Bot{
		ctx:      ctx,
		stop:     cancel,
		Token:    token,
		Updates:  make(chan Update, 50),
		handlers: make(map[string]HandlerFunc),
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
	if bot.Timeout == 0 {
		bot.Timeout = 5 * time.Second
	}

	me, err := bot.getMe()
	if err != nil {
		return nil, err
	}
	bot.Me = me

	return bot, nil
}

func (b *Bot) Start() {
	ctx, _ := context.WithCancel(b.ctx)
	go b.Hook.Update(ctx, b, b.Updates)
	for {
		select {
		case up := <-b.Updates:
			b.processUpdate(&up)
		case <-b.ctx.Done():
			log.Println("Waiting hook task finished...")
			time.Sleep(b.Timeout)
			return
		}
	}
}

func (b *Bot) Stop() {
	b.stop()
}

func (b *Bot) processUpdate(up *Update) {
	if up.Message == nil {
		return
	}
	msg := up.Message
	if msg.Text != "" {
		b.handle(OnText, msg)
	}
}

func (b *Bot) Send(id string, text string) {
	b.sendMessage(id, text)
}
