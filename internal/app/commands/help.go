package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, "You wrote: "+inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help - help\n"+"/list - list products")
	c.bot.Send(msg)
}
