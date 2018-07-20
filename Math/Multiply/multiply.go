package multiply

func Multiply321(multiplicand int) int {
	product := 0
	product += multiplicand << 8
	product += multiplicand << 6
	product += multiplicand
	return product
}

func Multiply321Add(multiplicand int) int {
	product := multiplicand
	product += product
	product += product
	product += product
	product += product
	product += product
	product += product

	product1 := product
	product1 += product1
	product1 += product1

	return multiplicand + product + product1
}
