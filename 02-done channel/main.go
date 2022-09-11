package main

import (
	"fmt"
	"runtime"
)

func generateData(done <-chan struct{}) <-chan int {
	data := make(chan int, 1)

	go func(data chan int) {
		defer close(data)

		for i := 0; i < 10; i++ {
			select {
			case <-done:
				return
			case data <- i:
			}
		}

	}(data)

	return data
}

func main() {
	done := make(chan struct{})
	data := generateData(done)

	counter := 0
	for d := range data {
		fmt.Println(d)

		counter++
		if counter == 5 {
			done <- struct{}{} // we want it done when counter is five.
		}
	}

	fmt.Printf("expected 1 goroutine, got goroutine: %d\n", runtime.NumGoroutine())
}
