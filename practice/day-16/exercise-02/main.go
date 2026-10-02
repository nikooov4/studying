package main

import "context"

func Double(
	ctx context.Context,
	in <-chan int,
) (<-chan int, <-chan struct{}) {

	out, finiched := make(chan int), make(chan struct{})

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(out)
				close(finiched)
				return
			case v, ok := <-in:
				if !ok {
					close(out)
					close(finiched)
					return
				}
				select {
				case <-ctx.Done():
					close(out)
					close(finiched)
					return
				case out <- v * 2:
				}

			}
		}
	}()
	return out, finiched
}

func main() {

}
