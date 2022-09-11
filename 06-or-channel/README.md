# Or Channel


Or-channel is used to listen multiple events, and wish to keep going after received done event from one of them.

For example, the program listen A, B, C three events, we want to notify other people after received only one done event from them.

Refer [Or Channel Example Code](./main.go)

Except `sync.Once` in the example to avoid close channel repeatedly, you could refer `orRecursive` this function to avoid close channel repeatedly. Personally, I like `sync.Once`.