package main

import "fmt"

func PlusFiveAdress(number *int) {
	*number += 5
}

func PlusFive(number int) {
	number += 5
}

func main() {
	ten := 10 // 10
	PlusFive(ten) // 15 но копия
	fmt.Println(ten) // 10
	PlusFiveAdress(&ten) //15 в оригинал записали
	fmt.Println(ten) // 15
	
}	
