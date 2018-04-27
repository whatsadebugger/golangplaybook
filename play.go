package main

import (
	"fmt"
	"reflect"
)

func main() {
	type T struct {
		AB int
		CD string
	}
	t := T{1337, "ahmad"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	s.Field(0).SetInt(2000)
	s.Field(1).SetString("Chickens")
	fmt.Println("t is now", t)
}
