package main

import (
	"fmt"
	"sync"
)

func incrementCounter(
	counter *int,
	wg *sync.WaitGroup,
	mx *sync.Mutex,
) {
	defer wg.Done()
	mx.Lock()
	*counter++
	mx.Unlock()
}

func main() {
	var counter int
	wg := &sync.WaitGroup{}
	mx := &sync.Mutex{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go incrementCounter(
			&counter,
			wg,
			mx,
		)
	}
	wg.Wait()
	fmt.Println(counter)
}
