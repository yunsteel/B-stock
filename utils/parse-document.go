package utils

import (
	"strings"

	"golang.org/x/net/html"
)

type Product struct {
	Name          string
	RegularPrice  string
	DiscountPrice string
	Link          string
	isSoldOut     bool
}

func parseDocument(document *html.Node) []Product {
	res := []Product{}

	var parse func(*html.Node)

	parse = func(n *html.Node) {
		if hasClass(n, "item-pay") && hasClass(n.Parent.Parent, "item-detail") {
			itemInfo := extractProductInfo(n)

			if n.Parent.Data == "a" {
				href := n.Parent.Attr[0].Val
				itemInfo.Link = getSiteURL() + href
			}

			if !itemInfo.isSoldOut {
				res = append(res, *itemInfo)
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parse(c)
		}
	}

	parse(document)
	return res
}

func extractProductInfo(n *html.Node) *Product {
	product := Product{}

	if n.Type == html.TextNode && strings.TrimSpace(n.Data) == "SOLDOUT" {
		product.isSoldOut = true
	}

	if n.Type == html.ElementNode {
		switch n.Data {
		case "h2":
			product.Name = strings.TrimSpace(n.FirstChild.Data)
		default:
			if hasClass(n, "sale_pay") {
				product.RegularPrice = strings.TrimSpace(n.FirstChild.Data)
			} else if hasClass(n, "pay") && hasClass(n, "inline-blocked") {
				product.DiscountPrice = strings.TrimSpace(n.FirstChild.Data)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		info := extractProductInfo(c)

		if product.Name == "" {
			product.Name = info.Name
		}
		if product.RegularPrice == "" {
			product.RegularPrice = info.RegularPrice
		}
		if product.DiscountPrice == "" {
			product.DiscountPrice = info.DiscountPrice
		}
		if product.Link == "" {
			product.Link = info.Link
		}
		if info.isSoldOut {
			product.isSoldOut = true
		}
	}

	return &product
}

func hasClass(n *html.Node, className string) bool {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "class" && strings.Contains(attr.Val, className) {
				return true
			}
		}
	}
	return false
}
