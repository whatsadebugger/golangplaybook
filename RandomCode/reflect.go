package main

import (
	"fmt"
	"reflect"
)

func main() {
	type ahmad struct {
		s string
	}

	wow := ahmad{"wow"}
	fmt.Println(reflect.TypeOf(wow))
	fmt.Println(reflect.ValueOf(wow).String())
}
