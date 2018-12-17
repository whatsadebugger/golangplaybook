package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

type testStructTags struct {
	Ahmad int `json:"ahmad,damha"`
}

func main() {
	JSONS := []string{
		`{"ahmad":1, "nest": {"wow":100,"look":"nohands","nested":{"key": 1337}}}`,
		`{"Count":1, "wow":"wow"}`,
		`{"damha":5}`,
		`{"ahmad": 0, "damha":5}`,
		`{"damha": 0, "ahmad":5}`,
	}

	for _, jason := range JSONS {
		var tst map[string]interface{}
		if err := json.Unmarshal([]byte(jason), &tst); err != nil {
			panic(err)
		}
		fmt.Println(tst)
		data, _ := json.Marshal(tst)
		fmt.Println(string(data))
	}

	jad := []byte(`{"Count":1, "Wow":"wow", "nest": {"key":1}}`)
	var test map[string]interface{}
	if err := json.Unmarshal(jad, &test); err != nil {
		panic(err)
	}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Wow}} and nest {{.nest.key}}\n")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(os.Stdout, test)
}
