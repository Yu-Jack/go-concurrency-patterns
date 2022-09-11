# Done Channel

Done channel is used as coordinator to communicate between different goroutine about which one is done. 

In the example, you could see I choose to close done channel to send message because this characteristic `Can read data from closed channel`.

We could also use `context.WithCancel()` to create context which could be called by `cancel()`. Then you could use `case <-ctx.Done()` to receive the cancellation message. In the following examples, I'll use cancelable context to avoid goroutine leaking.

Beside the done channel, there is also a timeout mechanism to avoid execution time is too long. Use `case <- time.After(10 * time.Seconds)` in `select` could receive message after 10 seconds, then you could do `return` to exit the goroutine. 

