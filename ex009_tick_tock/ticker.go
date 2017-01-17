package main

import (
	"time"
	"fmt"
)

func main() {

	ticks := myTick(1 * time.Second)

	go func() {
		for t := range ticks {
			fmt.Println("ticking at", t)
		}
	}()

	time.NewTicker()

	time.Sleep(5 * time.Second)
	fmt.Println("Done.")
}

func myTick(d time.Duration) (ticks chan time.Time) {

	ticks = make(chan time.Time)
	go func() {
		for {
			time.Sleep(d)
			ticks <- time.Now()
		}
	}()
	return
}


