package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	layTheTableWith("cheese", "ham", "eggs")

	commenceShovelling("Alice","Rita","Sue","Bob")

	<-tableIsCleared
	fmt.Println("that was delicious!")
}

func layTheTableWith(dishes ...string) {
	morsels := make([]string, 0)
	for _, dish := range dishes {
		amountOfMorsels := between(5, 10)
		for i := 0; i < amountOfMorsels; i++ {
			morsels = append(morsels, dish)
		}
	}

	fmt.Println(len(morsels), "morsels on the table")

	go func() {
		for _, v := range rand.Perm(len(morsels)) {
			morselChannel <- morsels[v]
		}

		tableIsCleared <- true
	}()
}

func commenceShovelling(people ...string) {
	for _, person := range people {
		go eat(person)
	}
}

func eat(name string) {
	for {
		morsel := <-morselChannel
		fmt.Println(name, "is enjoying some", morsel)
		timeToEat := time.Duration(between(3, 6)) * time.Second
		<-time.After(timeToEat)
	}
}

func between(low, hi int) int {
	return rand.Intn(hi - low + 1) + low
}

var morselChannel = make(chan string)
var tableIsCleared = make(chan bool)