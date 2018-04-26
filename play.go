package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 3
	fmt.Println("type:", reflect.TypeOf(x))
}
