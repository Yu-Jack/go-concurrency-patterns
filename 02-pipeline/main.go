package main

import (
	"fmt"
	"runtime"
)

func generateData() <-chan int {
	data := make(chan int, 1)

	go func() {
		defer close(data)

		for i := 0; i < 10; i++ {
			data <- i
		}
	}()

	return data
}

func add(input <-chan int, target int) <-chan int {
	added := make(chan int, 1)

	go func() {
		defer close(added)

		for data := range input {
			added <- data + target
		}
	}()

	return added
}

func multiply(input <-chan int, target int) <-chan int {
	added := make(chan int, 1)

	go func() {
		defer close(added)

		for data := range input {
			added <- data * target
		}
	}()

	return added
}

func main() {
	for d := range multiply(add(generateData(), 1), 2) {
		fmt.Println(d)
	}

	fmt.Printf("expected 1 goroutine, got goroutine: %d\n", runtime.NumGoroutine())
}
