package main

import (
	"fmt"

	"main.go/utils"
)

func main() {
	const MAX_AVAILABLE_PAGE = 5

	inStockItems := []string{}

	for i := 1; i <= MAX_AVAILABLE_PAGE; i++ {
		page := utils.Crawl(i)
		inStockItems = append(inStockItems, page...)
	}

	fmt.Println(inStockItems)
}
