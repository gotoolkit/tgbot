package main

import (
	"log"

	"github.com/gotoolkit/tgbot"
)

func main() {
	bot, err := tgbot.New("554451803:AAH97LZgnF4LmlpRFbp24Col82PLLxMyZhE")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(bot.Me)
	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	bot.Stop()
	// }()
	bot.Start()

}
