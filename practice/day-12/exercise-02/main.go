package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func incrementCounter(
	counter *atomic.Int64,
	n int,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		counter.Add(1)
	}
}

func main() {
	var counter atomic.Int64
	n := 1000
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go incrementCounter(&counter, n, wg)
	}
	wg.Wait()
	fmt.Println(counter.Load())
}
