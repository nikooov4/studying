package main

import "fmt"

func main() {
	num := [4]int{1, 2, 3, 4}
	s := num[:2] // 1 2 | 0 0
	t := s // 1 2 | 0 0
	t = append(t, 99) // 1 2 99 | 0

	fmt.Println(num, s, t, 
		len(num), cap(num), 
		len(s), cap(s), 
		len(t), cap(t))
}
