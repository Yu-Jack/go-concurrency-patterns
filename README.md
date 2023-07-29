# Introduction

This repo demos for golang concurrency patterns after read [Concurrency in Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/). There are something I modified in these examples.

1. Buffered or unbuffered channel both are okay, it doesn't have effect in these examples. But for memory efficiency, I choose unbuffered channel.
2. I reorder the pattern for easily learning based on my opinion. So there might be some differences between repo and book. 
3. In each example, it will print `expected 1 goroutine, got goroutine: 1`. That means the program close all goroutines safely (except main goroutine) to avoid goroutine leaks.

There is an simple introduction [article](https://yu-jack.github.io/2022/10/04/go-concurrency/) written in Chinese. Feel free to read it.
# Outline



| Concurrency Pattern                           |
|-----------------------------------------------|
 | [01: Generator](./01-generator)               | 
 | [02: Done channel](./02-done-channel)         | 
 | [03: Pipeline](./03-pipeline)                 | 
 | [04: Fan-In](./04-fan-in)                     | 
 | [05: Fan-Out](./05-fan-out)                   | 
 | [06: Fan-In and Fan-Out](./06-fan-in-fan-out) |
 | [07: Or-Channel](./07-or-channel)             |
 | [08: Or-Done-Channel](./08-or-done-channel)   |
 | [09: Tee-Channel](./09-tee-channel)           |
 | [10: Bridge-Channel](./10-bridge-channel)     |

# Related Usage

These examples are built with above related concepts. 

| Other Usage                                  |
|:---------------------------------------------|
| [01: Graceful Shutdown](./other/01-graceful) | 
| [02: Pub/Sub](./other/02-pub-sub)            |
| [03: Worker Pool](./other/03-worker-pool)    |

# References

1. [Concurrency in Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/). 
