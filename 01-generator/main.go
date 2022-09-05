package main

import "fmt"

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

func main() {
	data := generateData()

	for d := range data {
		fmt.Println(d)
	}
}
