package main

import (
	"fmt"
	"sync"
	"time"
)

func count(f, t int, result *int) {
	for f < t+1 {
		*result += f * f
		f++
	}
}

func separate(n int, wg *sync.WaitGroup) int {
	i := n / 2
	var sum1 int
	var sum2 int

	wg.Add(1)
	go func() {
		defer wg.Done()
		count(i+1, n, &sum2)
	}()
	count(1, i, &sum1)

	wg.Wait()

	return sum1 + sum2
}

func main() {
	var sum1 int
	var wg sync.WaitGroup

	start1 := time.Now()
	count(1, 2_000_000, &sum1)
	time1 := time.Since(start1)

	start2 := time.Now()
	sum2 := separate(2_000_000, &wg)
	time2 := time.Since(start2)

	fmt.Println(time1, time2, sum1, sum2)

}
