package main

import (
	"encoding/json"
	"fmt"
)

type testStructTags struct {
	Ahmad int `json:"ahmad,damha"`
}

func main() {
	JSONS := []string{`{"ahmad":1}`, `{"damha":5}`, `{"ahmad": 0, "damha":5}`, `{"damha": 0, "ahmad":5}`}

	for _, jason := range JSONS {
		var tst testStructTags
		if err := json.Unmarshal([]byte(jason), &tst); err != nil {
			panic(err)
		}
		fmt.Println(tst)
	}
}
