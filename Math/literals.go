package main

import "fmt"

func main() {

	fmt.Println("%d")
	fmt.Printf("%d\n\n", 0xBadFace)
	
	fmt.Println("%x")
	fmt.Printf("%x\n\n", 0xBadFace)
	
	fmt.Println("%E")
	fmt.Printf("%E\n", .1E17)
	fmt.Printf("%E\n", 6.67428E-11)
	fmt.Printf("%E\n\n", 1.e+0)

	fmt.Println("%f with modified precision")
	fmt.Printf("%.0f\n", .1E17)
	fmt.Printf("%.17f\n", 6.67428E-11)
	fmt.Printf("%.17f\n", 1.e+0)
}

