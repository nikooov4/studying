package main

import (
	"context"
	"slices"
	"testing"
	"time"
)

func equal(arr1, arr2 []int) bool {

	slices.Sort(arr1)
	slices.Sort(arr2)
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func TestMain1(t *testing.T) {

	jobs := make(chan int, 3)
	jobs <- 1
	jobs <- 2
	jobs <- 3

	count := 5
	ctx := context.Background()
	test := []int{2, 4, 6}
	close(jobs)
	output := Process(ctx, jobs, count)
	outputArr := []int{}

	for v := range output {
		outputArr = append(outputArr, v)
	}

	if !equal(test, outputArr) {
		t.Errorf("fatal")
	}

}

func TestMain2(t *testing.T) {

	jobs := make(chan int, 1)
	jobs <- 1

	count := 5
	ctx, cancel := context.WithCancel(context.Background())
	out := Process(ctx, jobs, count)
	cancel()
	for {
		select {
		case _, ok := <-out:
			if ok == false {
				return
			}
		case <-time.After(1 * time.Second):
			t.Fatal("fatal")
		}
	}
}
