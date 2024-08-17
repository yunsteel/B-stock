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

			filteredItems := utils.FilterItems(func(item string) bool {
				return strings.Contains(strings.ToLower(item), strings.ToLower(update.Message.Text))
			}, inStockItems)

			if len(filteredItems) == 0 {
				return
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, strings.Join(filteredItems, "\n"))
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
