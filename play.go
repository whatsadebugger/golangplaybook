package main

import "fmt"

func main() {

	attended := map[string]bool{
		"Ann": true,
		"Joe": true,
	}

	if attended["Ann"] { // will be false if person is not in the map
		fmt.Println("Ann", "was at the meeting")
	}
}
