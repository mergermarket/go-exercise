package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {

	freeComputers := makeComputers(12)

	for i := 0; i < 25; i++ {
		go touristShowsUp(fmt.Sprintf("Tourist %d", i), freeComputers)
	}

	fmt.Scanln()
}

func makeComputers(slots int) chan bool {

	freeComputers := make(chan bool, slots)

	for i := 0; i < slots; i++ {
		freeComputers <- true
	}

	return freeComputers
}

func touristShowsUp(name string, slots chan bool) {

	select {
	case <-slots:
		useComputer(name, slots)
	default:
		fmt.Println(name, "is waiting for turn.")
		<-slots
		useComputer(name, slots)
	}
}

func useComputer(name string, slots chan bool) {
	fmt.Println(name, "is online")
	waitPeriod := time.Duration(between(2,5)) * time.Second
	<-time.After(waitPeriod)
	fmt.Println(name, "is done, having spent", waitPeriod, "online")
	slots <- true
}

func between(low, hi int) int {
	return rand.Intn(hi - low + 1) + low
}
