package main

import (
	"time"
	"math/rand"
	"fmt"
)

func main() {

	go play("Ed");
	go play("Natasha");

	time.Sleep(30 * time.Second)
	fmt.Println("Time up!")
}

func play(name string) {
	score:= 0
	for {
		roll := rollDice()
		wait := 6-roll
		score += roll
		fmt.Printf("%s (%d) rolled a %d, waiting %d sec\n",
			name, score, roll, wait)

		time.Sleep(time.Duration(wait) * time.Second)
	}
}

func rollDice() int {
	return rand.Intn(5) + 1
}


