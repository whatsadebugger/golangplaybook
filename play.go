package main

import (
	"fmt"
)

// import

func main() {
	// add minutes and hours. If negative get final minutes and hours and add + 24 +60
	// if non negative get final minutes and hours and create object
	h, m := 1, -4820
	tom := h*60 + m
	fm := tom % 60
	fh := (tom % 24) + 24

	fmt.Println(fh, fm)
}
