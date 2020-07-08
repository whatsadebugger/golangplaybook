package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {

	chA, chB := make(chan string, 100), make(chan string, 100)

	wg.Add(3)
	go readFromFile(chA)
	go first(chA, chB)
	go second(chB)
	wg.Wait()
}

func readFromFile(chA chan string) {
	f, err := os.Open("./text.txt")
	noErr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		printSend("A", scanner.Text())
		chA <- scanner.Text()
	}
	close(chA)
	wg.Done()
}

func first(chA, chB chan string) {
	for v := range chA {
		printRead("A", v)
		printSend("B", v)
		chB <- v
	}
	close(chB)
	wg.Done()
}

func second(chB chan string) {
	for v := range chB {
		printRead("B", v)
		fmt.Println("Reversed word from channel B", reverse(v))
	}
	wg.Done()
}

func printRead(ch, word string) {
	fmt.Printf("Read from Channel %s the word %s\n", ch, word)
}

func printSend(ch, word string) {
	fmt.Printf("Sending %s to Channel %s\n", ch, word)
}

func reverse(text string) string {
	size := len(text)
	rev := make([]byte, size)

	for i, j := size-1, 0; i >= 0; i, j = i-1, j+1 {
		rev[j] = (byte)(text[i])
	}

	return string(rev)
}

func noErr(err error) {
	if err != nil {
		panic(err)
	}
}
