package main

import "fmt"

func main() {
	defer fmt.Println("outer")

	func() {
		defer fmt.Println("inner")
		fmt.Println("body")
	}()

	fmt.Println("after")
}
