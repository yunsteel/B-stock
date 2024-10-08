package utils

import "strings"

func Filter[T any](filter func(item T) bool, items []T) []T {
	res := []T{}

	for _, item := range items {
		if filter(item) {
			res = append(res, item)
		}
	}

	return res
}

func Map[T any, R any](m func(item T) R, items []T) []R {
	res := []R{}

	for _, item := range items {
		res = append(res, m(item))
	}

	return res
}

func SelectProductByKeyword(keyword string, items []Product) []Product {
	return Filter(func(product Product) bool {
		return strings.Contains(strings.ToLower(product.Name), strings.ToLower(keyword))
	}, items)
}
