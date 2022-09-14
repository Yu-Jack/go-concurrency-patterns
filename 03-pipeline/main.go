package main

import (
	"fmt"
	"runtime"
)

func generateData(done <-chan struct{}) <-chan int {
	data := make(chan int)

	go func() {
		defer close(data)

		////marked, please read readme
		// simulate hardworking
		//time.Sleep(200 * time.Millisecond)

		for i := 0; i < 10; i++ {
			select {
			case data <- i:
			case <-done:
				return
			}
		}
	}()

	return data
}

func add(done <-chan struct{}, input <-chan int, target int) <-chan int {
	added := make(chan int)

	go func() {
		defer close(added)

		for {
			select {
			case v, ok := <-input:
				if !ok {
					return
				}
				added <- v + target
			case <-done:
				return
			}
		}
	}()

	return added
}

func multiply(done <-chan struct{}, input <-chan int, target int) <-chan int {
	multiplied := make(chan int)

	go func() {
		defer close(multiplied)

		for {
			select {
			case v, ok := <-input:
				if !ok {
					return
				}
				multiplied <- v * target
			case <-done:
				return
			}
		}
	}()

	return multiplied
}

func main() {
	done := make(chan struct{})

	////marked, please read readme
	//time.AfterFunc(1000*time.Millisecond, func() {
	//	close(done)
	//})

	for d := range multiply(done, add(done, generateData(done), 1), 2) {
		fmt.Println(d)
	}

	fmt.Printf("expected 1 goroutine, got goroutine: %d\n", runtime.NumGoroutine())
}
