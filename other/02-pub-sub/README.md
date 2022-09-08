# Pub/Sub

這實現了非常簡易的 Pub/Sub 功能

參考 [Pub/Sub Example Code](./main.go)

另外以下的程式中註解那段可以打開，並把上面給註解掉，就會發現第一個 publish 就有接收到。理由很單純，就是時間差問題而已，這裡單純是展示說晚點 close channel 可以收到訊息而已。


https://github.com/Yu-Jack/go-concurrency-patterns/blob/68a0ef6eb90bb382726ea217f65f9cc418fdfb54/other/02-pub-sub/main.go#L118-L125