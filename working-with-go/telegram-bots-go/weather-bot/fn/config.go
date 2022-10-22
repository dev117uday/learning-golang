package fn

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// TelegramToken : token for tg
var TelegramToken string

// WeatherAPIKey :  token for weather API
var WeatherAPIKey string

// LoadConfig : to load config and return env variable
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	TelegramToken = os.Getenv("telegramToken")
	WeatherAPIKey = os.Getenv("WeatherAPIKey")

}
