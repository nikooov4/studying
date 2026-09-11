package main

import "fmt"

type Player struct {
	Score int 
}

func (p Player) AddValue() {
	p.Score += 5
}

func (p *Player) AddPointer() {
	p.Score += 5
}

func main() {
	p := Player{
		Score: 10,
	}
	p.AddValue() // +5
	fmt.Println(p.Score) // 10
	p.AddPointer() // +5
	fmt.Println(p.Score) // 15
}
