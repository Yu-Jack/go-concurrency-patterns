package main

import (
	"fmt"
	"runtime"
	"time"
)

func generateData() <-chan int {
	data := make(chan int, 1)

	go func(data chan int) {
		for i := 0; i < 10; i++ {
			data <- i
		}

		close(data)
	}(data)

	return data
}

func fanOut(data <-chan int) <-chan int {
	target := make(chan int, 1)

	go func(data <-chan int) {
		defer close(target)
		for val := range data {
			target <- val
		}
	}(data)

	return target
}

func main() {
	data := generateData()
	target1 := fanOut(data)
	target2 := fanOut(data)

	go func() {
		for d := range target1 {
			fmt.Printf("target1: %d\n", d)
		}
	}()

	go func() {
		for d := range target2 {
			fmt.Printf("target2: %d\n", d)
		}
	}()

	time.Sleep(1 * time.Second) // Easily use time.Sleep to wait them finished instead of sync.WaitGroup

	fmt.Printf("expected 1 goroutine, got goroutine: %d\n", runtime.NumGoroutine())
}
