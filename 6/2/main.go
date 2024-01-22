package main

import (
	"context"
	"fmt"
	"time"
)

// exitCtx is a function representing the work of a goroutine
func exitCtx(ctx context.Context) {
	fmt.Println("Goroutine starts")

	// use select to wait for the context to be canceled
	select {
	case <-ctx.Done():
		// if the context is canceled, print a shutdown message
		fmt.Println("Shutdown request")
	default:
		// if the context is not canceled, finish the goroutine
		fmt.Println("Goroutine finish")
	}
}

func main() {
	// create a context and a cancel function to control the goroutine
	ctx, cancel := context.WithCancel(context.Background())

	// launch a goroutine with the exitCtx function and pass the context to it
	go exitCtx(ctx)

	// wait for some time to let the goroutine execute
	time.Sleep(2 * time.Second)

	// cancel the context to terminate the goroutine
	cancel()

	// wait for some more time for the goroutine to finish
	time.Sleep(1 * time.Second)

	fmt.Println("Main Goroutine finish")
}
