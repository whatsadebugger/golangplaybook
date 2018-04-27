package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "wow a string"
	v := reflect.ValueOf(&str)
	fmt.Println("Can set: ", v.Elem().CanSet())
	v.Elem().SetString("WOW i changed it")

	fmt.Println("pointer to string: ", v.Interface())
	fmt.Println("modified string: ", str)
}
