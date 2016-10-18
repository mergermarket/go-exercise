package main

import (
	"time"
	"fmt"
	"sort"
)

func main() {
	amountOfPlayers := 10
	lengthOfGame := 30

	eventStream := make(chan GameEvent)
	players := makePlayers(amountOfPlayers, eventStream)
	finalScores := gameLoop(lengthOfGame, eventStream, players)

	printScores(finalScores)
}

func makePlayers(amountOfPlayers int, eventStream chan GameEvent) (players []Player) {
	for i := 0; i < amountOfPlayers; i++ {
		player := NewPlayer(fmt.Sprintf("Player %d", i + 1), eventStream)
		players = append(players, player)
	}
	return
}

func gameLoop(lengthOfGame int, gameEvents <- chan GameEvent, players []Player) (scores []Score) {
	runningTotals := make(map[string]int)
	timesUp := time.After(time.Duration(lengthOfGame) * time.Second)

	for _, p := range players {
		p.Start()
	}

	for {
		select {
		case event := <-gameEvents:
			runningTotals[event.player.Name] += event.score
			wait := 6 - event.score
			fmt.Printf("%s: %s rolled a %d, waiting %d seconds\n", time.Now(), event.player.Name, event.score, wait)
			event.player.RollAgainAfter(wait)
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

