package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 3
	fmt.Println("value:", reflect.ValueOf(x).String())
}
