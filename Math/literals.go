package main

import "fmt"

func main() {

	fmt.Println("%d")
	fmt.Printf("%d\n\n", 0xBadFace)

	fmt.Println("%x")
	fmt.Printf("%x\n\n", 0xBadFace)

	fmt.Println("%E")
	fmt.Printf("%E\n", .1E-17)

	fmt.Println("%f with modified precision")
	fmt.Printf("%.18f\n", .1E-17)

}
