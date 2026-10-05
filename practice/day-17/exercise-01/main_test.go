package main

import (
	"strings"
	"testing"
)

func TestDecodeCreateTask(t *testing.T) {

	reader := strings.NewReader(`{"title":"Read"}`)
	task, errCreate := DecodeCreateTask(reader)

	if errCreate != nil {
		t.Fatalf("error: %v", errCreate)
	}
	if task.Title != "Read" {
		t.Fatal("error: wrong json")
	}

	reader2 := strings.NewReader(`{}`)
	task2, errCreate2 := DecodeCreateTask(reader2)

	if errCreate2 != nil {
		t.Fatalf("error: %v", errCreate2)
	}
	if task2.Title != "" {
		t.Fatal("error: wrong json")
	}

	reader3 := strings.NewReader(`{"lol""pool"}`)
	_, errCreate3 := DecodeCreateTask(reader3)

	if errCreate3 == nil {
		t.Fatalf("error: %v", errCreate3)
	}

	reader4 := strings.NewReader(`{"title":123}`)
	_, errCreate4 := DecodeCreateTask(reader4)

	if errCreate4 == nil {
		t.Fatalf("error: %v", errCreate4)
	}   
}
