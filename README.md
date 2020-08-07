# queue
```go
import "github.com/smallnest/queue"
```
Package queue provides multiple queue implementations. 
The lock-free and two-lock algorithms are from Michael and Scott. https://doi.org/10.1145/248052.248106

## Queue

A queue is a collection of entities that are maintained in a sequence and can be modified by the addition of entities at one end of the sequence and the removal of entities from the other end of the sequence.The operations of a queue make it a first-in-first-out (FIFO) data structure. In a FIFO data structure, the first element added to the queue will be the first one to be removed. 

Package queue defines `Queue` interface:

```go
type Queue interface {
	Enqueue(v interface{})
	Dequeue() interface{}
}
```

Currently it contains three implementations:

- lock-free queue: `LKQueue`
- two-lock queue: `CQueue`
- slice-based queue: `SliceQueue`


## Benchmark


```sh
goos: darwin
goarch: amd64
pkg: github.com/smallnest/queue

BenchmarkQueue/lock-free_queue#4-4           	 8399941	       177 ns/op
BenchmarkQueue/two-lock_queue#4-4            	 7544263	       155 ns/op
BenchmarkQueue/slice-based_queue#4-4         	 6436875	       194 ns/op

BenchmarkQueue/lock-free_queue#32-4          	 8399769	       140 ns/op
BenchmarkQueue/two-lock_queue#32-4           	 7486357	       155 ns/op
BenchmarkQueue/slice-based_queue#32-4        	 4572828	       235 ns/op

BenchmarkQueue/lock-free_queue#1024-4        	 8418556	       140 ns/op
BenchmarkQueue/two-lock_queue#1024-4         	 7888488	       155 ns/op
BenchmarkQueue/slice-based_queue#1024-4      	 8902573	       218 ns/op
```