package main

import (
	"log"

	"github.com/neufeldtech/joe-example/pkg/bot"
)

func main() {

	config := bot.Config{
		HTTPListen: "0.0.0.0:8080",
	}

	b, err := bot.NewBot(config)
	if err != nil {
		log.Fatal(err)
	}

	b.Respond("open sesame", bot.OpenDoor)

	b.Brain.RegisterHandler(b.HandleDoorbell)

	err = b.Run()
	if err != nil {
		b.Logger.Fatal(err.Error())
	}
}
