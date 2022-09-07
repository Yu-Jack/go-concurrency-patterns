# Or-Done Channel

從前面的 or channel 會發現一件事情，萬一有一個 channel 永遠是 blocking 的話就會有 leak goroutine 的風險存在，所以要透過 or-done channel 去避免 leak goroutine 這件事情發生的機會。

不過正常來說 or-done channel 是針對兩個 channels 去做 or，這邊的實作做成**對多個 channel 做 or**，所以有一點不太一樣，詳細實作參考 [Or-Done Channel Example Code](./main.go)