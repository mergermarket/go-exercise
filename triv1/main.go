package main

import (
	"math/rand"
	"time"
	"fmt"
)

var tasks = make(chan bool)
var alarm = make(chan bool)

func main() {

	fmt.Println("Let's go for a walk!")

	do("getting ready")("Alice")
	do("getting ready")("Bob")

	<-tasks
	<-tasks

	armTheAlarm()

	do("putting on shoes")("Alice")
	do("putting on shoes")("Bob")

	<-tasks
	<-tasks

	fmt.Println("Exiting and locking the door")

	<-alarm
}

func do(task string) func(string) {
	return func(name string) {
		go func() {
			wait := rand.Intn(6) + 1
			fmt.Println(name, "started", task)
			<-time.After(time.Duration(wait) * time.Second)
			fmt.Println(name, "spent", wait, "seconds", task)
			tasks <- true

		}()
	}
}

func armTheAlarm() {
	go func() {
		fmt.Println("Arming the alarm.")
		fmt.Println("Alarm is counting down.")
		<-time.After(10 * time.Second)
		fmt.Println("Alarm is armed.")
		alarm <- true
	}()
}

