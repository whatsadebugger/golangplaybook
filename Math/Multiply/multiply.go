package multiply

func Multiply321(multiplicand int) int {
	product := 0
	product += multiplicand << 8
	product += multiplicand << 6
	product += multiplicand
	return product
}
