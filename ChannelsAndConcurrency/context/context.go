package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel() // undefer if you want to see what is logged in ctx.Err() when called

	sleepAndTalk(ctx, 3*time.Second, "hello")
}

func sleepAndTalk(ctx context.Context, t time.Duration, s string) {
	select {
	case <-time.After(t):
		fmt.Println(s)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
