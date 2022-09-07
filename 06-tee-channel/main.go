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

func tee(input <-chan int, outputs []chan int) {
	go func() {
		var wg sync.WaitGroup

		for data := range input {
			wg.Add(len(outputs))

			for i, _ := range outputs {
				go func(i int, data int) {
					defer wg.Done()
					time.Sleep(time.Duration(i) * time.Second)
					outputs[i] <- data
				}(i, data)
			}
		}

		wg.Wait()

		for _, channel := range outputs {
			close(channel)
		}
	}()
}

func main() {
	outputs := make([]chan int, 2)
	for i, _ := range outputs {
		outputs[i] = make(chan int)
	}

	input := generateData()
	tee(input, outputs)

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
