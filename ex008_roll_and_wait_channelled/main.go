package main

import (
	"time"
	"fmt"
	"math/rand"
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
	sort.Sort(sort.Reverse(ByFinalScore(finalScores)))
	fmt.Println(finalScores)
	fmt.Println(finalScores[0].Name, "is the WINNER!")
}

type GameEvent struct {
	player *Player
	score  int
}

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

type Score struct {
	Name  string
	Score int
}

type ByFinalScore []Score

func (items ByFinalScore) Len() int {
	return len(items)
}

func (items ByFinalScore) Less(x, y int) bool {
	return items[x].Score < items[y].Score
}

func (s ByFinalScore) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

