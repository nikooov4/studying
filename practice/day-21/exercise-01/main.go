package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {

	back := context.Background()
	ctrlC := os.Interrupt
	ctx, stop := signal.NotifyContext(back, ctrlC)
	defer stop()
	fmt.Println("waiting for signal")
	
	select {
	case <-ctx.Done():
		fmt.Println("signal received")
	}
}
