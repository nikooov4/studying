package main

import "context"

func StartSend(
	ctx context.Context,
	value int,
) (<-chan int, <-chan struct{}) {

	values := make(chan int)
	finished := make(chan struct{})

	go func() {
		defer close(finished)
		defer close(values)

		select {
		case values <- value:
		case <-ctx.Done():

		}

	}()

	return values, finished
}

func main() {

}
