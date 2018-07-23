package main

import (
    "fmt"
    "os"
    "os/signal"
    "time"
)

// https://chilts.org/2017/06/12/cancelling-multiple-goroutines
func main() {
    // a channel to tell `tick()` to stop, and one to tell us they've stopped
    stopChan := make(chan struct{})
    stoppedChan := make(chan struct{})
    go tick(stopChan, stoppedChan)

    // listen for C-c
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
    fmt.Println("main: received C-c - shutting down")

    // tell the goroutine to stop
    fmt.Println("main: telling goroutines to stop")
    close(stopChan)
    // and wait for them to reply back
    <-stoppedChan
    fmt.Println("main: goroutine has told us they've finished")
}

func tick(stop, stopped chan struct{}) {
    // tell the caller we've stopped
    defer close(stopped)

    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case now := <-ticker.C:
            fmt.Printf("tick: tick %s\n", now.UTC().Format(time.RFC1123))
        case <-stop:
            fmt.Println("tick: caller has told us to stop")
            return
        }
    }
}