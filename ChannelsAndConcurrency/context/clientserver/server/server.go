package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/whatsadebugger/golangplaybook/ChannelsAndConcurrency/context/clientserver/log"
)

func main() {
	http.HandleFunc("/", log.Decorate(handler))

	panic(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println(ctx, "handler started")
	defer log.Println(ctx, "handler finished")

	select {
	case <-time.After(time.Second * 3):
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
