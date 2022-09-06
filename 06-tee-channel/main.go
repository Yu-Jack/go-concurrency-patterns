package main

import (
	"fmt"
	"sync"
	"time"
)

func generateData() <-chan int {
	data := make(chan int, 1)

	go func(data chan int) {
		for i := 0; i < 2; i++ {
			data <- i
		}

		close(data)
	}(data)

	return data
}

func tee(input <-chan int, outputLen int) []chan int {
	channels := make([]chan int, outputLen)
	for i, _ := range channels {
		channels[i] = make(chan int)
	}

	go func() {
		var wg sync.WaitGroup

		for data := range input {
			wg.Add(outputLen)

			for i, _ := range channels {
				go func(i int, data int) {
					defer wg.Done()
					time.Sleep(time.Duration(i) * time.Second)
					channels[i] <- data
				}(i, data)
			}
		}

		wg.Wait()

		for _, channel := range channels {
			close(channel)
		}
	}()

	return channels
}

func main() {
	input := generateData()
	outputs := tee(input, 2)

	for i, output := range outputs {
		go func(output <-chan int, i int) {
			for out := range output {
				fmt.Printf("tee-%d: %d\n", i, out)
			}

			fmt.Printf("tee-%d: closed\n", i)
		}(output, i)
	}

	time.Sleep(5 * time.Second)
}
