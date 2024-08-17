package utils

import (
	"net/http"
	"strconv"

	"golang.org/x/net/html"
)

func CrawlPage(page int) []string {
	url := "https://www.swingguitars.com/612/?&page=" + strconv.Itoa(page) + "&sort=recent"
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	doc, err := html.Parse(res.Body)

	if err != nil {
		panic(err)
	}

	list := ParseDocument(doc)

	return list
}

func GetInStockItems() []string {
	const MAX_AVAILABLE_PAGE = 5

	inStockItems := []string{}

	for i := 1; i <= MAX_AVAILABLE_PAGE; i++ {
		page := CrawlPage(i)
		inStockItems = append(inStockItems, page...)
	}

	return inStockItems
}
