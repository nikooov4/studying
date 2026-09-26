package main

import (
	"fmt"
	"sync"
)

func incrementCounter(
	counter *int,
	n int,
	mx *sync.Mutex,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		mx.Lock()
		*counter++
		mx.Unlock()
	}
}

func main() {
	counter := 0
	n := 1000
	mx := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go incrementCounter(&counter, n, mx, wg)
	}
	wg.Wait()
	fmt.Println(counter)
}
