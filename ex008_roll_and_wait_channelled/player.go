package main

import (
	"math/rand"
)

type Player struct {
	Name  string
	ready chan bool
}

func NewPlayer(name string, game chan <- GameEvent) Player {
	ready := make(chan bool)
	player := Player{name, ready}

	go func() {
		for {
			<-ready
			roll := rollDice()
			game <- GameEvent{&player, roll}
		}
	}()

	return player
}

func (p *Player) Roll() {
	p.ready <- true
}

func rollDice() int {
	return rand.Intn(6) + 1
}

