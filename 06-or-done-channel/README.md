# Or-Done Channel


In previous example, there is a situation which it will block goroutine forever if one of channels never finish. It results in the goroutine leaks. So, we should use or-done channel to avoid goroutine leaking. In other hand, this example looks like timeout pattern in `select`.

Normal or-done channel version is for two channels, but here I implement for multiple channels. So, it's kind of different from normal version. Please refer [Or-Done Channel Example Code](./main.go).