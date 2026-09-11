package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4} // 1 2 3 4
	b := a // 1 2 3 4
	s := a[1:3] // 2 3
	b[0] = 99 // 99 2 3 4
	s[0] = 77 // 77 3

	fmt.Println(
		a, // 1 77 3 4
		b, // 99 2 3 4
		s, // 77 3
		len(s), // 2
		cap(s), // 3
	)
}
