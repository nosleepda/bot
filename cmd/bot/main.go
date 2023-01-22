package main

import (
	"github.com/joho/godotenv"
	"github.com/nosleepda/bot/internal/service/product"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	for update := range updates {
		if update.Message != nil { // If we got a message

			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
			case "list":
				listCommand(bot, update.Message, productService)
			default:
				defaultBehavior(bot, update.Message)
			}
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, "You wrote: "+inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help - help\n"+"/list - list products")
	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, productService *product.Service) {
	outputMessageText := "Here all of products: \n\n"
	products := productService.List()
	for _, p := range products {
		outputMessageText += p.Title
		outputMessageText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMessageText)
	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

	bot.Send(msg)
}
