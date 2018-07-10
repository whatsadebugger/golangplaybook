package main

import (
	"math/big"
	"fmt"
)

func main() {
	fmt.Println(9223372036854775807)
	b := big.NewInt(1e0)
	fmt.Println(b.Add(b,b).Text(10))
	fmt.Println(b)
}