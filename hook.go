package tgbot

import (
	"context"
	"log"
	"time"
)

type Hook interface {
	Update(context.Context, *Bot)
	Stop()
}

type UpdateHook struct {
}

var DefaultHook = &UpdateHook{}

func (u *UpdateHook) Update(ctx context.Context, b *Bot) {
	for {
		select {
		case <-ctx.Done():
			log.Println("done")
			return
		default:
			log.Println("sleep")
			time.Sleep(time.Second)
		}
	}
}
