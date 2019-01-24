package main

import (
	"fmt"
	"github.com/gotoolkit/tgbot"
	"log"
	"strconv"
)

func main() {
	bot, err := tgbot.New("554451803:AAH97LZgnF4LmlpRFbp24Col82PLLxMyZhE")
	if err != nil {
		log.Fatal(err)
	}
	bot.DebugFunc = func(up *tgbot.Update) {
		log.Println(up.Message)
	}
	bot.Handle(tgbot.OnText, func(m *tgbot.Message) {
		fmt.Println(m.Text)
		//bot.Send(strconv.Itoa(m.From.ID), m.Text)
	})

	bot.Handle("/ping", func(m *tgbot.Message) {
		bot.Send(strconv.Itoa(m.From.ID), "pong")
	})
	bot.Start()

}
