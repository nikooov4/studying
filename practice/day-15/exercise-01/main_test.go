package main

import (
	"context"
	"testing"
	"time"
)

func TestMain1(t *testing.T) {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		100*time.Millisecond,
	)
	defer cancel()
	ch := make(chan string)
	str, err := waitForResult(ctx, ch)
	if str != "" || err != context.DeadlineExceeded {
		t.Errorf("fatal")
	}

}

func TestMain2(t *testing.T) {

	ctx := context.Background()
	ch := make(chan string, 1)
	ch <- "ok"
	str, err := waitForResult(ctx, ch)
	if str != "ok" || err != nil {
		t.Errorf("fatal")
	}

}
