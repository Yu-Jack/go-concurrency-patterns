package main

import (
	"context"
	"fmt"
	"runtime"
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

func orDone(ctx context.Context, sources ...<-chan int) <-chan int {
	var once sync.Once
	target := make(chan int)

	for _, source := range sources {
		go func(target chan int, source <-chan int) {
			defer once.Do(func() {
				close(target)
			})

			for {
				select {
				case v, ok := <-source:
					if !ok {
						return
					}

					target <- v
				case <-ctx.Done():
				}
			}

		}(target, source)
	}

	return target
}

func tee(ctx context.Context, input <-chan int, outputs []chan int) {
	go func() {
		var wg sync.WaitGroup

		for data := range orDone(ctx, input) {
			wg.Add(len(outputs))

			for _, output := range outputs {
				go func(data int, output chan int) {
					defer wg.Done()

					select {
					case output <- data:
					case <-ctx.Done():
					}
				}(data, output)
			}
		}

		wg.Wait()

		for _, output := range outputs {
			close(output)
		}
	}()
}

func main() {
	outputs := make([]chan int, 2)
	for i, _ := range outputs {
		outputs[i] = make(chan int)
	}

	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(500*time.Millisecond, cancel)

	input := generateData()
	tee(ctx, input, outputs)

	for i, output := range outputs {
		go func(output <-chan int, i int) {
			for out := range output {
				fmt.Printf("tee-%d: %d\n", i, out)
			}

			fmt.Printf("tee-%d: closed\n", i)
		}(output, i)
	}

	time.Sleep(2 * time.Second)

	fmt.Printf("expected 1 goroutine, got goroutine: %d\n", runtime.NumGoroutine())
}
