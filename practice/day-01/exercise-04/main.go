package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5} // 1 2 3 4 5
	s := arr[1:4] // 2 3 4 
	s = s[:1] // 2 
	s = s[:3] // 2 3 4
	s[2] = 99 // 2 3 99

	fmt.Println(
		arr, // 1 2 3 99 5
		s, // 2 3 99
		len(s), // 3
		cap(s), // 4
	)
}
