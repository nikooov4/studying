package main

import "fmt"

func Redirect(one, two *int) {
	one = two // one (&address) = two (&address)
	*one = 99 // one (value) = 99 -> two = 99 (the same address)
}

func main() {
	x := 10
	y := 20
	p := &x 

	Redirect(p, &y) // 10 20
	fmt.Println(x, y, *p) // 10 99 99
}
