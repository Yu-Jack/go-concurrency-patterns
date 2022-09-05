# Generator


通常可以拿來做 pipeline 使用，避免一次讀取太大量資料。 其他使用方式還有類似 graceful shutdown 的方法。

參考 [Generator Pattern](./main.go)

補充: 通常產生 channel 人最好要負責關閉 channel，原因有以下兩點

1. 無法向已 closed channel 寫入資料
2. 可以讀取已 closed channel 的資料