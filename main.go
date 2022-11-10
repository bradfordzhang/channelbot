package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	tgbotToken := os.Getenv("BOT_TOKEN")

	if tgbotToken == "" {
		log.Printf("No BOT_TOKEN Found! Exiting..")
		os.Exit(1)
	}

	bot, err := tgbotapi.NewBotAPI(tgbotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if strings.Contains(update.Message.Text, "官网") {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "网址： https://www.fishport.cyou\n点击右侧 \"Buy it now\" 即可")
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}

			if strings.Contains(update.Message.Text, "截图") {
				msg := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL("https://raw.githubusercontent.com/bradfordzhang/res/main/%E6%88%AA%E5%9B%BE.jpg"))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}

			if strings.Contains(update.Message.Text, "钓鱼") {
				msg := tgbotapi.NewVideo(update.Message.Chat.ID, tgbotapi.FileURL("https://raw.githubusercontent.com/bradfordzhang/res/main/%E9%92%93%E9%B1%BC.mp4"))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
		}
	}
}
