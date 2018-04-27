package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string `json:"name"`
	Msg  string `json:"message"`
}

func main() {

	m := Message{Name: "ahmad", Msg: "ahmad is learning so much everyday"}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
	var m2 Message
	json.Unmarshal(b, &m2)
	fmt.Printf("Name: %s\nMessage: %s", m2.Name, m2.Msg)

}
