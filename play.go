package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf("Wow a string")

	fmt.Println(v.Interface())
}
