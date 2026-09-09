package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3} // 1 2 3
	s := arr[:] // 1 2 3
	t := s // 1 2 3
	
	t = append(t, 99) // 1 2 3 99 
	t[0] = 77 // 77 2 3 99

	fmt.Println(arr, // 1 2 3 
		s, // 1 2 3
		t, // 77 2 3 99
	len(s), // 3
	len(t)) // 4
}
