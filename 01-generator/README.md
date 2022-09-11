# Generator

Generator is function that returns a channel. 

Refer [Generator Pattern Example Code](./main.go)

Note: The generator which create channel should be responsible for closing its channel because there are two characteristics.

1. Can't send data into closed channel
2. Can read data from closed channel


It could be used as pipeline to avoid reading a lot of data or graceful shutdown, you could see [01: Graceful Shutdown](/other/01-graceful).