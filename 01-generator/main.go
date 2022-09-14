package main

import (
	"fmt"
	"runtime"
)

func generateData() <-chan int {
	data := make(chan int)

	go func(data chan int) {
		for i := 0; i < 10; i++ {
			data <- i
		}

		close(data)
	}(data)

	return data
}

func main() {
	data := generateData()

	for d := range data {
		fmt.Println(d)
	}

	fmt.Printf("expected 1 goroutine, got goroutine: %d\n", runtime.NumGoroutine())
}
