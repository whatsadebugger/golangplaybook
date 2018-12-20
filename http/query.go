package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {

	u, err := url.Parse("https://www.google.com/")
	if err != nil {
		panic(err)
	}

	fmt.Println("url:", u)
	fmt.Println("url path:", u.Path)
	fmt.Println("url rawpath:", u.RawPath)
	fmt.Println("url string:", u.String())

	u, _ = u.Parse("v1/v1/v1")
	fmt.Println("url:", u)
	fmt.Println("url path:", u.Path)
	fmt.Println("url rawpath:", u.RawPath)
	fmt.Println("url string:", u.String())

	r, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic(err)
	}

	q := url.Values{}
	q.Add("key1", "val1")
	q.Add("key2", "val2")
	r.URL.RawQuery = q.Encode()

	fmt.Println(r.URL.String())
}
