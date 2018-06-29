package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	address := `"0x6cc5f688a315f3dc28a7781717a9a798a59fda7b"`
	JSON := `{ "address": ` + address + `, "amount": 10 }`

	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &jsonMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonMap["address"])
	fmt.Println(jsonMap["amount"])
}
