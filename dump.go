package main

import (
	"encoding/json"
	"fmt"
)

func printMap(space string, m map[string]interface{}) {
	for k, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			fmt.Printf("\"%v\": \n{\n", k)
			printMap(space+"\t", mv)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v %v : %v\n", space, k, v)
		}
	}
}

var jsonStr = `
{
  "array": [
	1,
	2,
	3
  ],
  "boolean": true,
  "null": null,
  "number": 123,
  "object": {
	"a": "b",
	"c": "d",
	"e": "f"
  },
  "string": "Hello World"
}
`

func main() {
	jsonMap := make(map[string]interface{})
	if err := json.Unmarshal([]byte(jsonStr), &jsonMap); err != nil {
		panic(err)
	}
	printMap("", jsonMap)
}
