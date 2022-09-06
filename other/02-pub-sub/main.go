package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var subscribers = make(map[string][]chan int)
var mtx sync.Mutex
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func subscribe(name string) (<-chan int, string) {
	mtx.Lock()
	defer mtx.Unlock()

	ch := make(chan int)
	subscribers[name] = append(subscribers[name], ch)

	return ch, RandStringRunes(5)
}

func publish(name string, msg int) {
	mtx.Lock()
	defer mtx.Unlock()

	if chs, ok := subscribers[name]; ok {
		for _, ch := range chs {
			go func(ch chan int) {
				ch <- msg
			}(ch)
		}
	}
}

func printChannel(ch <-chan int, name string) {
	for data := range ch {
		fmt.Printf("subscriber-%s: %d\n", name, data)
	}
}

func main() {
	s1, name1 := subscribe("name-1")
	s2, name2 := subscribe("name-1")
	s3, name3 := subscribe("name-2")

	go printChannel(s1, name1)
	go printChannel(s2, name2)
	go printChannel(s3, name3)

	publish("name-1", 1)
	publish("name-1", 2)
	publish("name-2", 3)

	time.Sleep(1 * time.Second)
}
