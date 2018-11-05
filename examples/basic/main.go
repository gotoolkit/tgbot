package main

import (
	"log"
	"strconv"

	"github.com/gotoolkit/tgbot"
)

func main() {
	bot, err := tgbot.New("554451803:AAH97LZgnF4LmlpRFbp24Col82PLLxMyZhE")
	if err != nil {
		log.Fatal(err)
	}
	bot.Handle(tgbot.OnText, func(m *tgbot.Message) {
		bot.Send(strconv.Itoa(m.From.ID), m.Text)
	})
	bot.Start()

}
