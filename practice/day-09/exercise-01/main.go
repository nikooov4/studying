package main

import "fmt"

func Producer(n int, ch chan int) {
	i := 1
	for i <= n {
		ch <- i
		i++
	}

}

func Consumer(ch, summ chan int, n int) {
	sum := 0
	i := 0
	for i < n {
		sum += <-ch
		i++
	}
	summ <- sum
}

func main() {
	// ch1 := make(chan int, 0)
	// ch2 := make(chan int, 1)
	ch3 := make(chan int, 10)

	// sum1 := make(chan int)
	// sum2 := make(chan int)
	sum3 := make(chan int)

	// go Producer(1_000_000, ch1)
	// go Consumer(ch1, sum1, 1_000_000)

	// go Producer(1_000_000, ch2)
	// go Consumer(ch2, sum2, 1_000_000)

	go Producer(1_000_000, ch3)
	go Consumer(ch3, sum3, 1_000_000)

	// fmt.Println(len(ch1), <-sum1)
	// fmt.Println(cap(ch2), 1_000_000, <-sum2)
	fmt.Println(cap(ch3), 1_000_000, <-sum3)
}
