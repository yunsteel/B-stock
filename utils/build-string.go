package utils

import "strings"

func BuildProductString(product Product) string {
	상품명 := "상품명: " + product.Name
	가격 := "가격: " + product.DiscountPrice + "원 (정상가 " + product.RegularPrice + "원)"
	링크 := product.Link

	res := []string{상품명, 가격, 링크}

	return strings.Join(res, "\n")

}
