package main

import (
	"fmt"
	"math/big"
)

func main() {
	b := big.NewInt(1e0)
	fmt.Println(b, "initial")
	fmt.Println(b.Add(b, b).String(), "double yourself")
	fmt.Println(b.Mul(b, big.NewInt(1000)).String(), "Multiply by 1000")
}
