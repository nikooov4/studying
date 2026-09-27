package main

import (
	"fmt"
	"sync"
)

func Producer1(ch chan int) {
	defer close(ch)

	for i := 1; i <= 3; i++ {
		ch <- i
	}

}

func Producer2(ch chan int) {

	defer close(ch)
	ch <- 10
	ch <- 20

}

func Forwarding1(
	result, ch1 chan int,
	wg *sync.WaitGroup,
) {

	defer wg.Done()

	for v := range ch1 {
		result <- v
	}
}

func Forwarding2(
	result, ch2 chan int,
	wg *sync.WaitGroup,
) {

	defer wg.Done()

	for v := range ch2 {
		result <- v
	}
}

func main() {

	wgFor := &sync.WaitGroup{}

	ch1 := make(chan int)
	ch2 := make(chan int)
	result := make(chan int)

	go Producer1(ch1)
	go Producer2(ch2)
	wgFor.Add(1)
	go Forwarding1(result, ch1, wgFor)
	wgFor.Add(1)
	go Forwarding2(result, ch2, wgFor)
	go func() {
		wgFor.Wait()
		close(result)
	}()

	for v := range result {
		fmt.Println(v)
	}
}
