package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrond args", args)
		return
	}

	product, err := c.productService.Get(arg)
	if err != nil {
		log.Fatal(err)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.Title)
	c.bot.Send(msg)
}
