package main

import (
	"fmt"
	"runtime"
)

func generateData(done <-chan struct{}) <-chan int {
	data := make(chan int, 1)

	go func(data chan int) {
		defer close(data)

		i := 0
		for {
			select {
			case <-done:
				return
			case data <- i:
				i++
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
		fmt.Println(counter, d)
		counter++
		if counter == 5 {
			close(done)
			break
		}
	}

	fmt.Printf("expected 1 goroutine, got goroutine: %d\n", runtime.NumGoroutine())
}
