package main

import (
	"log"
	"os"

	"github.com/anonyindian/gottbot"
	"github.com/anonyindian/gottbot/ext"
	"github.com/anonyindian/gottbot/filters"
	"github.com/anonyindian/gottbot/handlers"
)

func main() {
	bot, err := gottbot.NewBot(os.Getenv("TAMTAM_BOT_TOKEN"), nil)
	if err != nil {
		panic(err)
	}
	updater := ext.NewUpdater(nil)
	updater.StartPolling(bot, nil)

	dispatcher := updater.Dispatcher

	dispatcher.AddHandler(handlers.MessageHandler(filters.Message.All, echo))

	updater.Idle()
}

func echo(bot *gottbot.Bot, update *gottbot.Update) error {
	msg := update.Message
	_, err := msg.Reply(bot, msg.Body.Text, nil)
	if err != nil {
		log.Println("failed to send message:", err.Error())
	}
	return ext.EndGroups
}