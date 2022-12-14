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
		log.Fatal("No BOT_TOKEN Found! Exiting..")
	}

	bot, err := tgbotapi.NewBotAPI(tgbotToken)
	if err != nil {
		log.Panic(err)
	}

	if os.Getenv("DEBUG") == "" {
		bot.Debug = false
		log.Println("Debug off")
	} else {
		bot.Debug = true
		log.Println("Debug on")
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			switch update.Message.Command() {
			case "start":
			case "about":
				msg.Text = "I'm zyc's channel bot.\nSource code: https://github.com/bradfordzhang/channelbot/"
			case "status":
				msg.Text = "I'm ok!"
			}

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			continue
		}

		if strings.Contains(update.Message.Text, "截图") {
			msg := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL("https://raw.githubusercontent.com/bradfordzhang/res/main/%E6%88%AA%E5%9B%BE.jpg"))
			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

		if strings.Contains(update.Message.Text, "钓鱼") {
			msg := tgbotapi.NewVideo(update.Message.Chat.ID, tgbotapi.FileURL("https://raw.githubusercontent.com/bradfordzhang/res/main/%E9%92%93%E9%B1%BC.mp4"))
			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

		/*if strings.Contains(update.Message.Text, "机场") {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "自用机场推荐: https://s.zyc.name/AUIoQWV/")
			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}*/
	}
}
