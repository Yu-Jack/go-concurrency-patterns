# Pipeline

Pipeline is composited by many Generator. 

There is a `marked` comment in example, you could try to uncomment it, you will see the program exit after 1 second because it has closed done channel.

This demos when generator doesn't have `i < 10` condition, and we hope there is no any goroutine leaks. So, remember to add done channel or cancelable context into goroutine creator to avoid the goroutine leaks.

Refer [Pipeline Example Code](./main.go)