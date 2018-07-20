package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("reflect.TypeOf")
	fmt.Println(reflect.TypeOf(0xBADFACE))
	fmt.Println(reflect.TypeOf(.1E-17))
	fmt.Println(reflect.TypeOf(0333))

	fmt.Println()
	fmt.Println()

	fmt.Println("%d")
	fmt.Printf("%d\n\n", 0xBadFace)
	fmt.Println("%d")
	fmt.Printf("%d\n\n", 0xBadFace)

	fmt.Println("%x")
	fmt.Printf("%x\n\n", 0xBadFace)

	fmt.Println("%E")
	fmt.Printf("%E\n\n", .1E-17)

	fmt.Println("%f with modified precision")
	fmt.Printf("%.18f\n", .1E-17)

}
