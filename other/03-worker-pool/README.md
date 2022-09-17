# Worker Pool

It implements the worker pool. I combine generator and fan-in pattern to achieve it. Besides, I add a lot of `<-ctx.Done()` to make sure all goroutine will be closed, except main goroutine.

Refer [Worker Pool Example Code](./main.go).