package telegram

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"main.go/utils"
)

func SendMessage() {
	utils.LoadDotEnv()
	token := os.Getenv("TELEGRAM_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	inStockItems := utils.GetInStockItems()

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			selectedItems := utils.SelectProductByKeyword(update.Message.Text, inStockItems)

			if len(selectedItems) == 0 {
				return
			}

			items := utils.Map(utils.BuildProductString, selectedItems)

			message := tgbotapi.NewMessage(update.Message.Chat.ID, strings.Join(items, "\n\n"))
			message.ReplyToMessageID = update.Message.MessageID

			bot.Send(message)
		}
	}
}
