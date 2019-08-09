package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/telegram-bot-api.v4"

	"github.com/fatih/color"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_API_KEY"))
	if err != nil {
		log.Fatalf("Error in instantiating the bot: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		color.Red("Error while receiving messages: %s", err)
		return
	}
	color.Green("Started %s successfully", bot.Self.UserName)

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand()  {
      chatID := update.Message.Chat.ID
      username := update.Message.From.UserName
			msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Hello %s!!",username))
			bot.Send(msg)
		}
	}
}

