package main

import (
    "fmt"
    "os"
    "os/signal"
    "sync"
    "time"
)

func main() {
    // a channel to tell `tick()` and `tock()` to stop
    stopChan := make(chan struct{})

    // a WaitGroup for the goroutines to tell us they've stopped
    wg := sync.WaitGroup{}

    // a channel for `tick()` to tell us they've stopped
    wg.Add(1)
    go tick(stopChan, &wg)

    // a channel for `tock()` to tell us they've stopped
    wg.Add(1)
    go tick(stopChan, &wg)

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
    fmt.Println("main: received C-c - shutting down")

    fmt.Println("main: telling goroutines to stop")
    close(stopChan)

    wg.Wait()
    fmt.Println("main: all goroutines have told us they've finished")
}

func tick(stop chan struct{}, wg *sync.WaitGroup) {
    // tell the caller we've stopped
    defer wg.Done()

    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case now := <-ticker.C:
            fmt.Printf("tick: tick %s\n", now.UTC().Format(time.Stamp))
        case <-stop:
            fmt.Println("tick: caller has told us to stop")
            return
        }
    }
}