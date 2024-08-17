package utils

func BuildProductString(product Product) string {
	상품명 := "상품명: " + product.Name
	가격 := "가격: " + product.Discount_Price + "원 (정상가 " + product.Regular_Price + "원)"

	return 상품명 + "\n" + 가격

}
