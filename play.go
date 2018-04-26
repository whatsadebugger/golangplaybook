package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf("Wow a string").Interface()

	fmt.Println(v)
}
