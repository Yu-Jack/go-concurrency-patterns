package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func sleep(after time.Duration) <-chan int {
	c := make(chan int)

	time.AfterFunc(after, func() {
		close(c)
	})

	return c
}

func or(sources ...<-chan int) <-chan int {
	var once sync.Once
	target := make(chan int)

	for _, source := range sources {
		go func(target chan int, source <-chan int) {
			defer once.Do(func() {
				close(target)
				fmt.Println("close should only happen once")
			})

			select {
			case <-source:
			}
		}(target, source)
	}

	return target
}

func orRecursive(sources ...<-chan int) <-chan int {
	// this is slice length, not the channel size.
	if len(sources) == 0 {
		return nil
	}

	if len(sources) == 1 {
		return sources[0]
	}

	target := make(chan int)

	go func() {
		defer func() {
			close(target)
			fmt.Println("close might happen more than once because recursive")
		}()

		if len(sources) == 2 {
			select {
			case <-sources[0]:
			case <-sources[1]:
			}
			return
		}

		select {
		case <-sources[0]:
		case <-sources[1]:
		case <-orRecursive(sources[2:]...):
		}

	}()

	return target
}

func main() {
	start := time.Now()

	<-or(
		sleep(3*time.Second),
		sleep(1*time.Second),
		sleep(4*time.Second),
		sleep(2*time.Second),
	)

	fmt.Printf("done after %v\n", time.Since(start))

	time.Sleep(5 * time.Second) // avoid no any error occurring

	fmt.Printf("expected 1 goroutine, got goroutine: %d\n", runtime.NumGoroutine())
}
