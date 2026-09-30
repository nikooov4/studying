package main

import (
	"context"
	"sync"
)

func Process(
	ctx context.Context,
	jobs <-chan int,
	workerCount int,
) <-chan int {

	result := make(chan int)

	if workerCount <= 0 {
		close(result)
		return result
	}

	var wg sync.WaitGroup
	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()

			for {
				select {
				case v, ok := <-jobs:
					if !ok {
						return
					} else {
						select {
						case result <- v * 2:
						case <-ctx.Done():
							return
						}
					}
				case <-ctx.Done():
					return 
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func main() {

}
