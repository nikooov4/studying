package main

import "fmt"

type Player struct {
	Score int 
}

func (p *Player) Experiment() {
	pp := Player {
		Score: 50,
	}
	p = &pp
	p.Score = 99
	fmt.Println(p.Score) // 99
}

func main() {
	p := Player {
		Score: 10,
	}
	p.Experiment()
	fmt.Println(p.Score) // 10
}
