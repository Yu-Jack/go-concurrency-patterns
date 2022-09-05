# Graceful Shutdown

通常應用程式在接受到 `kill -9` `kill -15` 等等相關訊號會直接關閉，但是實際上可以透過程式去處理取得這些 `signal` 時，要怎麼安全地關閉應用程式。


參考 [Graceful Shutdown Example Code](./main.go)

https://github.com/Yu-Jack/go-concurrency-patterns/blob/54f045e707fc514270433e955ad25f52d41e6869/other/01-graceful/main.go#L40-L41

在範例程式中，可以試著換成 `newNormalContext`，並且在印出 0 的時候去中斷程式，會看到使用 `newNormalContext` 時並不會看到 1 被印出來，就意外的被中斷。

但是當使用 `newGracefulContext` 接收 signal 的情況下，可以做一些後續收尾的動作。所以用這個 `newGracefulContext` 看到 0 的時候去中斷程式，會看到 1 被印出來。