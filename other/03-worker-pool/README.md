# Worker Pool

It implements the worker pool. I combine generator and fan-in pattern to achieve it. Besides, I add a lot of `<-ctx.Done()` to make sure all goroutine will be closed, except main goroutine.

Refer [Worker Pool Example Code](./main.go).

https://github.com/Yu-Jack/go-concurrency-patterns/blob/fcce2193b57338a35226dfd58401ac5c42773388/other/03-worker-pool/main.go#L101-L122

You could call `close()` early by replacing L105 with `time.AfterFunc(2*time.Second, cancel)`. You still see the `got goroutine 1`.
