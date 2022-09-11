package main

import (
	"fmt"
	"runtime"
	"sync"
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

func main() {
	data1 := generateData()
	data2 := generateData()
	target := fanIn(data1, data2)

	for d := range target {
		fmt.Println(d)
	}

	fmt.Printf("expected 1 goroutine, got goroutine: %d\n", runtime.NumGoroutine())
}
