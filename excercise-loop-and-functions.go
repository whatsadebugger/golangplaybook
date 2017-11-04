package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z			:= 1.0
	

	for {
		zee := z - (z*z - x) / (2*z)
		if math.Abs((z*z) - x )< 0.00000000001 {
			return zee
		}
		z = zee
	}
}

func main() {
	fmt.Println(Sqrt(82))	
}
