package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	WithCancel()
	// WithTimeout()
}

func WithCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	go Gen(ctx)
	time.Sleep(time.Second)
	cancel()

	time.Sleep(time.Second)
}

func WithTimeout() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	go Gen(ctx)
	time.Sleep(time.Second * 2)
}

func Gen(ctx context.Context) {
	done := ctx.Done()
	if done == nil {
		fmt.Println("not cancelable")
		return
	}

	tv := time.NewTicker(time.Millisecond * 100)
	for {
		select {
		case <-done:
			fmt.Println("canceled")
			return
		case <-tv.C:
			fmt.Println("Producing")
		}
	}
}
