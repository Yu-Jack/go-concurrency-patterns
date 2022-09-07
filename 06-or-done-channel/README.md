# Or-Done Channel

從前面的 or channel 會發現一件事情，萬一有一個 channel 永遠是 blocking 的話就會有 leak goroutine 的風險存在，所以要透過 or-done channel 去避免 leah goroutine 這件事情發生的機會。

參考 [Or-Done Channel Example Code](./main.go)