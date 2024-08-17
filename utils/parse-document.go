package utils

import (
	"strings"

	"golang.org/x/net/html"
)

type Product struct {
	Name           string
	Regular_Price  string
	Discount_Price string
	Is_Sold_Out    bool
}

func ParseDocument(document *html.Node) []Product {
	res := []Product{}

	var parse func(*html.Node)

	parse = func(n *html.Node) {
		if hasClass(n, "item-pay") && hasClass(n.Parent.Parent, "item-detail") {
			itemInfo := ExtractProductInfo(n)
			if !itemInfo.Is_Sold_Out {
				res = append(res, ExtractProductInfo(n))
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parse(c)
		}
	}

	parse(document)
	return res
}

func ExtractProductInfo(n *html.Node) Product {

	product := Product{}

	if n.Type == html.TextNode && strings.TrimSpace(n.Data) == "SOLDOUT" {
		product.Is_Sold_Out = true
	}

	if n.Type == html.ElementNode {
		switch n.Data {
		case "h2":
			product.Name = strings.TrimSpace(n.FirstChild.Data)
		default:
			if hasClass(n, "sale_pay") {
				product.Regular_Price = strings.TrimSpace(n.FirstChild.Data)
			} else if hasClass(n, "pay") && hasClass(n, "inline-blocked") {
				product.Discount_Price = strings.TrimSpace(n.FirstChild.Data)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		info := ExtractProductInfo(c)

		if product.Name == "" {
			product.Name = info.Name
		}
		if product.Regular_Price == "" {
			product.Regular_Price = info.Regular_Price
		}
		if product.Discount_Price == "" {
			product.Discount_Price = info.Discount_Price
		}
		if info.Is_Sold_Out {
			product.Is_Sold_Out = true
		}
	}

	return product
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
