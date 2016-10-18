package main

import (
	"time"
	"math/rand"
	"fmt"
)

type Score struct {
	name  string
	score int
}

func main() {

	rolls := make(chan Score)
	timeout := time.After(5 * time.Second)
	scores := make(map[string]int)

	for i:=1; i<50; i++ {
		go play(fmt.Sprintf("Player %d",i), rolls)
	}

	Loop:
	for {
		select {
		case <-timeout:
			fmt.Println("Time's up")
			break Loop
		case roll := <-rolls:
			scores[roll.name] += roll.score
		}
	}

	fmt.Println(scores)
}

func play(name string, rolls chan <-Score) {
	for {
		roll := rollDice()
		wait := 6 - roll
		fmt.Printf("%s rolled a %d, waiting %d sec\n",
			name, roll, wait)
		rolls <- Score{name, roll}
		time.Sleep(time.Duration(wait) * time.Second)
	}
}

func rollDice() int {
	return rand.Intn(5) + 1
}


