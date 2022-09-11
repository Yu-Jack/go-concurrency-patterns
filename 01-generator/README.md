# Generator

Generator is used as pipeline to avoid reading a lot of data.

Refer [Generator Pattern Example Code](./main.go)

Note: The generator which create channel should be responsible for closing its channel because there are two characteristics.

1. Can't send data into closed channel
2. Can read data from closed channel


Other usage is like graceful shutdown, you could see [01: Graceful Shutdown](./other/01-graceful).