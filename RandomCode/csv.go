package main

import (
	"encoding/csv"
	"os"
)

func main() {

	type user struct {
		UserName string
		Password string
	}

	userbase := []user{
		user{
			UserName: "ahmad",
			Password: "12354",
		},
		user{
			UserName: "wow21",
			Password: "fjfjfada",
		},
	}

	// import some users first
	f, err := os.Open("users.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		userbase = append(userbase, user{UserName: line[0], Password: line[1]})
	}

	wr := csv.NewWriter(os.Stdout)

	for _, v := range userbase {
		if err := wr.Write([]string{v.UserName, v.Password}); err != nil {
			return
		}
	}

	wr.Flush()

}
