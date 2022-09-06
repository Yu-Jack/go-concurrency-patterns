# Or Channel

通常用在想要監聽多的完成的事件時可以利用此 Pattern，舉例來說可能當 A B C 其中一個事件完成，就要馬上通知別人。

參考 [Or Channel Example Code](./main.go)

除了範例使用 `sync.Once` 去預防避免重複關閉之外，可以參考 `orRecursive` 的方式去實現，以個人來說我比較喜歡 `sync.Once`