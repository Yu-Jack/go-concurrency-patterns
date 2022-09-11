# Tee Channel

Tee channel is used to spread the message to everywhere. For example, you have message called A. You could send A to person1 to do product and person2 to do minus with tee chanel.

This patterns is same concept with [Linux `tee` command](https://www.runoob.com/linux/linux-comm-tee.html). Although linux tee command only send data to two outputs, here I implement to send data to **multiple outputs**. Please refer [Tee Channel Example Code](./main.go)。

By the way, it's similar to Pub/Sub, you could check the very simple version [02-Pub/Sub](/other/02-pub-sub) in this repo.