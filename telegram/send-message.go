package telegram

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"main.go/utils"
)

var bot *tgbotapi.BotAPI

func loadTelegramBot() {
	utils.LoadDotEnv()
	token := os.Getenv("TELEGRAM_TOKEN")

	tgbot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Fatal(err)
	}

	tgbot.Debug = true

	log.Printf("Authorized on account %s", tgbot.Self.UserName)

	bot = tgbot
}

func sendMessage(update tgbotapi.Update, inStockItems []utils.Product) {
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

func Run() {
	loadTelegramBot()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	inStockItems := utils.GetInStockItems()

	for update := range updates {
		if update.Message != nil {
			sendMessage(update, inStockItems)
		}
	}
}
