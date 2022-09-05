package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func newNormalContext() context.Context {
	return context.Background()
}

func newGracefulContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sig := make(chan os.Signal)
		signal.Notify(sig) // accept all signal
		defer signal.Stop(sig)

		select {
		case <-ctx.Done():
		case <-sig:
			cancel()
		}

		fmt.Println("graceful shutdown")
	}()

	return ctx
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	ctx := newGracefulContext()
	//ctx := newNormalContext()
	go func(ctx context.Context) {
		defer wg.Done()

		data := 0
		fmt.Println(data)

		// simulate the hard work, then set data of 1
		time.Sleep(2 * time.Second)
		data = 1
		fmt.Println(data)

		select {
		case <-ctx.Done():
			fmt.Println("skip following job")
			return
		default:
		}

		time.Sleep(2 * time.Second)
		data = 2
		fmt.Println(data)
	}(ctx)

	wg.Wait()
}
