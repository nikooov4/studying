package main

import "fmt"

func main() {
	src := []int{1, 2, 3} // 1 2 3
	dst := make([]int, 2, 5) // 0 0 | 0 0 0

	n := copy(dst, src) // 2 
	//dst = 1 2 | 0 0 0
	src[0] = 99 
	//src = 99 2 3

	fmt.Println(dst, src, len(dst), cap(dst), n) // 1 2 / 99 2 3 //
	// 2, 5, 2
}
