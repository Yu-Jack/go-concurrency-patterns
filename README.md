# Introduction

This repo demos for golang concurrent patterns. 

# Outline

In each example, it will print `expected 1 goroutine, got goroutine: 1`. That means the program close all goroutines safely (except main goroutine) to avoid goroutine leaking.

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

| Other Usage                                  |
|:---------------------------------------------|
| [01: Graceful Shutdown](./other/01-graceful) | 
| [02: Pub/Sub](./other/02-pub-sub)            |