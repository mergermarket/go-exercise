package main

import (
	"time"
	"fmt"
	"sort"
)

func main() {
	amountOfPlayers := 10
	lengthOfGame := 5

	eventStream := make(chan GameEvent)

	makePlayers(amountOfPlayers, eventStream)

	finalScores := gameLoop(lengthOfGame, eventStream)

	printScores(finalScores)
}

func makePlayers(amountOfPlayers int, eventStream chan GameEvent) {
	for i := 0; i < amountOfPlayers; i++ {
		p := NewPlayer(fmt.Sprintf("Player %d", i + 1), eventStream)
		p.Roll()
	}
}

func gameLoop(lengthOfGame int, gameEvents <- chan GameEvent) (scores []Score) {
	runningTotals := make(map[string]int)
	timesUp := time.After(time.Duration(lengthOfGame) * time.Second)
	for {
		select {
		case event := <-gameEvents:
			runningTotals[event.player.Name] += event.score
			wait := 6 - event.score
			fmt.Println(event.player.Name, "rolled a", event.score, ", waiting", wait, "seconds", time.Now())
			go func() {
				<-time.After(time.Duration(wait) * time.Second)
				event.player.Roll()
			}()
		case <-timesUp:
			fmt.Println("Time's up!")
			for key, val := range runningTotals {
				scores = append(scores, Score{key, val})
			}
			return scores
		}
	}
}

func printScores(finalScores []Score) {
	sort.Sort(sort.Reverse(ByScore(finalScores)))
	fmt.Println(finalScores)
	fmt.Println(finalScores[0].Name, "is the WINNER!")
}

type GameEvent struct {
	player *Player
	score  int
}

