# Tee Channel

通常在想把同樣一個訊息送到各個地方去處理時可以使用此 Pattern，例如說拿這 A 訊息，想先送給 Person1 做加法，以及要送給 Person2 做乘法等等。

此 Pattern 跟 [Linux `tee` command](https://www.runoob.com/linux/linux-comm-tee.html) 是很相似的概念。雖然 linux `tee` command 是同時輸出到兩個地方，但這邊的實作是**同時輸出到多個地方**，所以有一點不一樣 ，詳細實作參考 [Tee Channel Example Code](./main.go)。

另外這類似 Pub/Sub 的功能，可以參考非常簡單版本的 [02-Pub/Sub](/other/02-pub-sub)。