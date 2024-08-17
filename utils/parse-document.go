package utils

import (
	"strings"

	"golang.org/x/net/html"
)

func ParseDocument(document *html.Node) []string {
	res := []string{}

	var parse func(*html.Node)

	parse = func(n *html.Node) {
		if hasClass(n, "item-pay") && hasClass(n.Parent.Parent, "item-detail") {
			itemInfo := ExtractItemInfo(n)
			if !strings.HasSuffix(strings.ToUpper(itemInfo), "SOLDOUT") {
				res = append(res, ExtractItemInfo(n))
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parse(c)
		}
	}

	parse(document)
	return res
}

func ExtractItemInfo(n *html.Node) string {
	var text string

	if n.Type == html.TextNode {
		text = n.Data
	} else if n.Type == html.ElementNode {
		switch n.Data {
		case "br":
			text += "\n"
		case "h2":
			text += "상품명:"
		default:
			if hasClass(n, "sale_pay") {
				text += "정상가:"
			} else if hasClass(n, "pay") && hasClass(n, "inline-blocked") {
				text += "판매가:"
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += " " + ExtractItemInfo(c)
	}

	return strings.TrimSpace(text)
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
