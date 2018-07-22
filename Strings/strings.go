package main

import (
	"fmt"
)

func main() {
	const placeOfInterest = `⌘`
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Printf("placeOfInterest string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted placeOfInterest: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	fmt.Printf("% X", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("sample string: ")
	fmt.Printf("%s", sample)
	fmt.Printf("\n")

	fmt.Printf("quoted placeOfInterest: ")
	fmt.Printf("%+q", sample)
	fmt.Printf("\n")

	fmt.Printf("Hex values in sample: % X", sample)
}
