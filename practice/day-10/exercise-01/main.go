package main

import (
	"fmt"
	"sync"
)

func generator(n int, ch chan int, wg *sync.WaitGroup) {
	i := 1
	for i <= n {
		ch <- i
		i++
	}
	defer close(ch)
	defer wg.Done()
}

func square(in chan int, out chan int, wg *sync.WaitGroup) {
	for el := range in {
		out <- el * el
	}
	defer close(out)
	defer wg.Done()
}

func consumer(in chan int, wg *sync.WaitGroup) {
	sum := 0
	for el := range in {
		fmt.Println(el)
		sum += el
	}
	fmt.Println(sum)
	defer wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg.Add(3)
	go generator(5, ch1, wg)
	go square(ch1, ch2, wg)
	go consumer(ch2, wg)
	wg.Wait()
}
