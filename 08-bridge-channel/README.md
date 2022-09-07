# Bridge Channel

這個 Pattern 跟 Fan-In 很相似，差別在與 Bridge 有要求 `順序性`。Fan-in 會發現是不按照順序送資料過來，但在一些想要保證 channel 順序性的情況下，就會需要用 bridge channel 去處理這件事情。

詳細實作參考 [Bridge Channel Example Code](./main.go)