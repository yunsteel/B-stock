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
		if isProductDetailNode(n) {
			product := extractProduct(n)

			if n.Parent.Data == "a" {
				href := n.Parent.Attr[0].Val
				product.Link = getSiteURL() + href
			}

			if !product.isSoldOut {
				res = append(res, *product)
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parse(c)
		}
	}

	parse(document)
	return res
}

func extractProduct(n *html.Node) *Product {
	product := Product{}

	if isSoldOut(n) {
		product.isSoldOut = true
	}

	if n.Type == html.ElementNode {
		if n.Data == "h2" {
			product.Name = strings.TrimSpace(n.FirstChild.Data)
		}

		if isRegularPriceNode(n) {
			product.RegularPrice = strings.TrimSpace(n.FirstChild.Data)

		} else if isDiscountPriceNode(n) {
			product.DiscountPrice = strings.TrimSpace(n.FirstChild.Data)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		info := extractProduct(c)

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

func isSoldOut(n *html.Node) bool {
	return n.Type == html.TextNode && strings.TrimSpace(n.Data) == "SOLDOUT"
}

func isProductDetailNode(n *html.Node) bool {
	return hasClass(n, "item-pay") && hasClass(n.Parent.Parent, "item-detail")
}

func isRegularPriceNode(n *html.Node) bool {
	return hasClass(n, "sale_pay")
}

func isDiscountPriceNode(n *html.Node) bool {
	return hasClass(n, "pay") && hasClass(n, "inline-blocked")
}
