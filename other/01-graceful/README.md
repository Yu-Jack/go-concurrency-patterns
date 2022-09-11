# Graceful Shutdown

After the application accept `kill -9` `kill -15`, it will be closed. Actually, program can control the shutdown flow by intercepting the signal from os, then you could shut down application safely.

Refer [Graceful Shutdown Example Code](./main.go).

https://github.com/Yu-Jack/go-concurrency-patterns/blob/54f045e707fc514270433e955ad25f52d41e6869/other/01-graceful/main.go#L40-L41

In the example, you could use `newNormalContext`. Then shutdown program (ctrl+c) while the console prints `0`. The console doesn't show `1` because program is closed accidentally.

But when use `newGracefulContext`, the console will print `1`. Because it intercepts the signal, then program can prepare to shutdown in advance to avoid being closed accidentally.