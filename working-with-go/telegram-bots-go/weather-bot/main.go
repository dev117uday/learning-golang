package main

import (
	"log"
	fn "telegramBots/weather-bot/fn"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	fn.LoadConfig()
	bot, err := tgbotapi.NewBotAPI(fn.TelegramToken)
	if err != nil {
		log.Fatal("Unable to connect to Telegram API", err)
		return
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	config := tgbotapi.UpdateConfig{}
	updates, err := bot.GetUpdatesChan(config)

	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		data, err := fn.GetWeather(update.Message.Text)
		dataMsg := fn.ToString(data)

		if err != nil {
			log.Println(err)

			errorMsg := "Unable to find the City. Check the spelling."
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, errorMsg)
			_, err = bot.Send(msg)
			if err != nil {
				log.Println("Unable to send message:", err)
				return
			}

			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, dataMsg)
		_, err = bot.Send(msg)
		if err != nil {
			log.Println("Unable to send message:", err)
			return
		}
		data = nil
	}

}
