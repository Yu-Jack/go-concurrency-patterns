package main

import (
	"context"
	"fmt"
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

func orDone(ctx context.Context, sources ...<-chan int) <-chan int {
	var once sync.Once
	target := make(chan int)

	for _, source := range sources {
		go func(target chan int, source <-chan int) {
			defer once.Do(func() {
				close(target)
				fmt.Println("close should only happen once")
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

func main() {
	start := time.Now()
	ctx, cancel := context.WithCancel(context.Background())

	// Trigger `close()` on purpose to see all channel is closed immediately.
	time.AfterFunc(500*time.Millisecond, cancel)

	<-orDone(
		ctx,
		sleep(3*time.Second),
		sleep(1*time.Second),
		sleep(4*time.Second),
		sleep(2*time.Second),
	)

	fmt.Printf("done after %v\n", time.Since(start))

	time.Sleep(1 * time.Second) // avoid no any error occurring
}
