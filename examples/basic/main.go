package main

import (
	"log"
	"net/http"

	"github.com/gotoolkit/tgbot"
)

func main() {
	bot, err := tgbot.New("554451803:AAFsqzrvk-FjzvOS2-1-_OsoUNuJ_6ygqTc", tgbot.WithClient(http.DefaultClient))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(bot.Me)
}
