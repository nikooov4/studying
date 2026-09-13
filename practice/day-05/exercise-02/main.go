package main

import "fmt"

func main() {
	x := 1

	defer fmt.Println("A", x)
	defer func() {
		fmt.Println("B", x)
	}()

	x = 2
	fmt.Println("C", x)
}
