package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type subscriber struct {
	name  string
	ch    chan int
	topic string
	close chan struct{}
}

var subscribers = make(map[string][]subscriber)
var mtx sync.Mutex
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func subscribe(topic string) subscriber {
	mtx.Lock()
	defer mtx.Unlock()

	sub := subscriber{
		name:  randStringRunes(5),
		ch:    make(chan int),
		topic: topic,
		close: make(chan struct{}, 1),
	}

	subscribers[topic] = append(subscribers[topic], sub)

	return sub
}

func tee(ctx context.Context, input int, outputs []subscriber) {
	for _, output := range outputs {

		go func(output subscriber) {
			select {
			case <-output.close:
				fmt.Printf("channe-%s-%s is closed, you can't send message into it.\n", output.topic, output.name)
				return
			default:
			}

			select {
			case output.ch <- input:
			case <-ctx.Done():
			}
		}(output)
	}
}

func publish(ctx context.Context, topic string, msg int) {
	mtx.Lock()
	defer mtx.Unlock()

	if subs, ok := subscribers[topic]; ok {
		if len(subs) != 0 {
			tee(ctx, msg, subs)
		}
	}
}

func unsubscribe(topic string, sub subscriber) {
	mtx.Lock()
	defer mtx.Unlock()

	if subs, ok := subscribers[topic]; ok {
		for i, s := range subs {
			if s.name == sub.name {
				close(s.ch)
				s.close <- struct{}{}

				newSubs := subs[:i]
				newSubs = append(newSubs, subs[i+1:]...)

				subscribers[topic] = newSubs
				return
			}
		}
	}
}

func printChannel(sub subscriber) {
	for data := range sub.ch {
		fmt.Printf("subscriber-%s-%s: %d\n", sub.topic, sub.name, data)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s1 := subscribe("name-1")
	s2 := subscribe("name-1")
	s3 := subscribe("name-2")

	go printChannel(s1)
	go printChannel(s2)
	go printChannel(s3)

	publish(ctx, "name-1", 1)
	publish(ctx, "name-1", 2)

	publish(ctx, "name-2", 3) // not print, because below unsubscribe close the channel before this publish
	unsubscribe("name-2", s3)
	publish(ctx, "name-2", 333) // not print

	//publish(ctx, "name-2", 3) // print
	//time.Sleep(500 * time.Millisecond)
	//unsubscribe("name-2", s3)
	//publish(ctx, "name-2", 333) // not print

	time.Sleep(5 * time.Second)

	fmt.Printf("expected 3 goroutine, got goroutine: %d\n", runtime.NumGoroutine())
}
