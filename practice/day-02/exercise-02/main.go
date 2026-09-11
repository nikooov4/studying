package main

import "fmt"

func Swap(one, two *int) {
	var three int
	three = *one
	*one = *two
	*two = three
}

func main() {
	one, two := 1, 2 // 1 2
	fmt.Println(one, two, &one, &two) // 1 2 
	Swap(&one, &two)
	fmt.Println(one, two, &one, &two) // 2 1
}
