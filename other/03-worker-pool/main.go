package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

const (
	workerSize = 3
	jobSize    = 20
)

func worker(ctx context.Context, id int, job <-chan int) <-chan int {
	result := make(chan int)

	go func() {
		defer close(result)

		for {
			select {
			case <-ctx.Done():
				return
			case j, ok := <-job:
				if !ok {
					return
				}

				log.Printf("worker id: %d, start to deal with job: %d", id, j)
				time.Sleep(1 * time.Second) // simulate the hard work

				select {
				// make sure if result channel is not writable during cancel(), we still could exit in this goroutine
				case <-ctx.Done():
					return
				case result <- j * 2:
				}

				log.Printf("worker id: %d, end to deal with job: %d", id, j)
			}
		}
	}()

	return result
}

func genJobs(ctx context.Context) <-chan int {
	jobs := make(chan int)

	go func() {
		defer close(jobs)

		for i := 0; i < jobSize; i++ {
			select {
			case <-ctx.Done():
				return
			case jobs <- i:
			}
		}
	}()

	return jobs
}

func fanIn(ctx context.Context, sources ...<-chan int) <-chan int {
	target := make(chan int)

	go func() {
		defer close(target)
		wg := sync.WaitGroup{}

		for _, source := range sources {
			wg.Add(1)
			go func(source <-chan int, target chan int) {
				defer wg.Done()

				for {
					select {
					case <-ctx.Done():
						return
					case val, ok := <-source:
						if !ok {
							return
						}
						target <- val
					}
				}
			}(source, target)
		}

		wg.Wait()
	}()

	return target
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := genJobs(ctx)
	results := make([]<-chan int, workerSize)

	for i := 0; i < workerSize; i++ {
		results[i] = worker(ctx, i, jobs)
	}

	flow := fanIn(ctx, results...)
	for data := range flow {
		log.Printf("Result is :%d", data)
	}

	fmt.Printf("expected 1 goroutine, got goroutine: %d\n", runtime.NumGoroutine())
}
