package tgbot

import (
	"context"
	"log"
	"time"
)

type Hook interface {
	Update(context.Context, *Bot, chan Update)
}

type UpdateHook struct {
	timeout      time.Duration
	lastUpdateID int
}

var DefaultHook = &UpdateHook{
	timeout: 3 * time.Second,
}

func (u *UpdateHook) Update(ctx context.Context, b *Bot, event chan Update) {
	ticker := time.NewTicker(time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			log.Println("hook stopped")
			return
		case <-ticker.C:
		}

		updates, err := b.getUpdates(u.lastUpdateID+1, u.timeout)
		if err != nil {
			log.Println(err)
			log.Println("Failed to getUpdates, retrying in 3 seconds...")
			time.Sleep(time.Second * 3)
			continue
		}

		for _, update := range updates {
			u.lastUpdateID = update.ID
			event <- update
		}
	}
}
