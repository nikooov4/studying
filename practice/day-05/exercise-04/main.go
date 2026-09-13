package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)
		}
	}()

	defer fmt.Println("cleanup")

	fmt.Println("before")
	panic("boom")
	fmt.Println("after")
}
