package main

import (
	"context"
	"testing"
)

func TestStartSendDeliversValue(t *testing.T) {

	ctx := context.Background()
	value := 5
	values, finished := StartSend(ctx, value)

	got := <-values
	if got != value {
		t.Fatalf("got: %v, expected: %v", got, value)
	}

	_ = <-finished

}

func TestStartSendStopsOnCancelWithoutConsumer(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	values, finished := StartSend(ctx, 5)
	cancel()

	_ = <-finished

	if _, ok := <-values; ok {
		t.Fatalf("values must be closed")
	}
}
