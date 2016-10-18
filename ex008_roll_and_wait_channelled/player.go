package main

import (
	"math/rand"
	"time"
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

func (p *Player) Start() {
	p.ready <- true
}

func (p *Player) RollAgainAfter(seconds int) {
	go func () {
		<-time.After(time.Duration(seconds) * time.Second)
		p.ready <- true
	}()
}

func rollDice() int {
	return rand.Intn(6) + 1
}

