package main

import (
	"fmt"
	"sync"
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

func fanIn(sources ...<-chan int) <-chan int {
	target := make(chan int, 1)

	go func() {
		defer close(target)

		wg := sync.WaitGroup{}

		for _, source := range sources {
			wg.Add(1)
			go func(source <-chan int, target chan int) {
				defer wg.Done()
				for val := range source {
					target <- val
				}
			}(source, target)
		}

		wg.Wait()

	}()

	return target
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
	data1 := generateData()
	data2 := generateData()

	target := fanIn(data1, data2)

	target1 := fanOut(target)
	target2 := fanOut(target)

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
}
