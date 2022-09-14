package main

import (
	"context"
	"fmt"
	"runtime"
)

func generateStreams() <-chan <-chan int {
	streams := make(chan (<-chan int))

	go func() {
		defer close(streams)

		for i := 0; i < 3; i++ {
			stream := generateData()
			streams <- stream
		}
	}()

	return streams
}

func generateData() <-chan int {
	data := make(chan int)

	go func(data chan int) {
		defer close(data)

		for i := 0; i < 4; i++ {
			data <- i
		}
	}(data)

	return data
}

func bridge(ctx context.Context, chanSteam <-chan <-chan int) <-chan int {
	b := make(chan int)

	go func() {
		defer close(b)

		for {
			var ch <-chan int

			select {
			case v, ok := <-chanSteam:
				if !ok {
					return
				}
				ch = v
			case <-ctx.Done():
				return
			}

			for val := range ch {
				select {
				case b <- val:
				case <-ctx.Done():
				}
			}
		}
	}()

	return b
}

func main() {
	inputs := generateStreams()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for val := range bridge(ctx, inputs) {
		fmt.Println(val)
	}

	fmt.Printf("expected 1 goroutine, got goroutine: %d\n", runtime.NumGoroutine())
}
