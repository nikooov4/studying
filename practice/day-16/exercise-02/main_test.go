package main

import (
	"context"
	"testing"
	"time"
)

func TestNormalFlow(t *testing.T) {
	in := make(chan int, 3)
	in <- 1
	in <- 2
	in <- 3
	close(in)
	ctx := context.Background()

	out, finished := Double(ctx, in)

	for i := 0; i < 3; i++ {
		select {
		case v, ok := <-out:
			if !ok {
				t.Fatal("out channel closed unexpectedly")
			}
			if v != (i+1)*2 {
				t.Fatalf("expected %d, got %d", (i+1)*2, v)
			}
		case <-finished:
			t.Fatal("finished channel closed unexpectedly")
		}
	}
	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("out channel should be closed")
		}
	case <-time.After(1 * time.Second):
		t.Fatal("timeout waiting for out channel to close")
	}
	select {
	case _, ok := <-finished:
		if ok {
			t.Fatal("finished channel should be closed")
		}
	case <-time.After(1 * time.Second):
		t.Fatal("timeout waiting for finished channel to close")
	}
}

func TestCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	in := make(chan int)
	sent := make(chan struct{})

	out, finished := Double(ctx, in)

	go func() {
		in <- 1
		close(sent)
	}()
	select {
	case <-sent:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timeout sent")
	}
	cancel()
	select {
	case <-finished:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timeout finished")
	}
	select {
	case _, ok := <-out:
		if ok {
			cancel()
			t.Fatal("fatal: out is open")
			return
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timeout out")
	}
}
