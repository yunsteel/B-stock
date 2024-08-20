package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"main.go/utils"
)

func Run(keyword string) {
	utils.LoadDotEnv()
	token := os.Getenv("TELEGRAM_TOKEN")

	res, err := http.Get("https://api.telegram.org/bot" + token + "/getUpdates")

	if err != nil {
		log.Panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}

	var update Update
	json.Unmarshal(body, &update)

	if !update.Ok {
		return
	}

	chatIds := []string{}

	for _, res := range update.Result {
		chatIds = append(chatIds, strconv.Itoa(res.Message.Chat.ID))
	}

	selectedItems := utils.Filter(func(item utils.Product) bool {
		return strings.Contains(strings.ToUpper(item.Name), strings.ToUpper(keyword))
	}, utils.GetInStockItems())

	items := utils.Map(func(product utils.Product) string {
		return utils.BuildProductString(product)
	}, selectedItems)

	if len(selectedItems) == 0 {
		fmt.Println(keyword + "에 해당하는 제품이 아직 없어요.")
		return
	}

	for _, id := range chatIds {
		baseURL := "https://api.telegram.org/bot" + token + "/sendMessage" + "?chat_id=" + id

		for _, item := range items {

			res, err := http.Get(baseURL + "&text=" + item)

			if err != nil {
				log.Panic(err)
			}

			fmt.Println(res)
		}
	}
}
