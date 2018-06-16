package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	// Scan for next token.
	for scanner.Scan() {
		fmt.Println("word: ", scanner.Text())
	}

}
