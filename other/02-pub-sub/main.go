package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type subscriber struct {
	name string
	ch   chan int
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
		name: randStringRunes(5),
		ch:   make(chan int),
	}

	subscribers[topic] = append(subscribers[topic], sub)

	return sub
}

func publish(topic string, msg int) {
	mtx.Lock()
	defer mtx.Unlock()

	if subs, ok := subscribers[topic]; ok {
		for _, sub := range subs {
			go func(sub subscriber) {
				sub.ch <- msg
			}(sub)
		}
	}
}

func unsubscribe(topic string, sub subscriber) {
	mtx.Lock()
	defer mtx.Unlock()

	if subs, ok := subscribers[topic]; ok {
		for i, s := range subs {
			if s.name == sub.name {
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
		fmt.Printf("subscriber-%s: %d\n", sub.name, data)
	}
}

func main() {
	s1 := subscribe("name-1")
	s2 := subscribe("name-1")
	s3 := subscribe("name-2")

	go printChannel(s1)
	go printChannel(s2)
	go printChannel(s3)

	publish("name-1", 1)
	publish("name-1", 2)
	publish("name-2", 3)

	unsubscribe("name-2", s3)

	publish("name-2", 333) // not print

	time.Sleep(1 * time.Second)
}
