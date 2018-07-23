package main

import (
	"log"
	"sync"
	"time"
)

// taken from https://pocketgophers.com/limit-concurrent-use/
func main() {
	log.SetFlags(log.Ltime)

	tasks := make(chan int)
	var wg sync.WaitGroup
	for worker := 0; worker < 3; worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for i := range tasks {
				time.Sleep(time.Second)
				log.Println(i)
			}
		}()
	}

	for i := 0; i < 9; i++ {
		tasks <- i
	}

	close(tasks)
	
	wg.Wait()
}