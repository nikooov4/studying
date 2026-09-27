package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	result <- 42
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)
	wg.Add(1)
	defer wg.Wait()
	go worker(ch, wg)

	select {
	case v := <-ch:
		fmt.Println(v)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	}
}
