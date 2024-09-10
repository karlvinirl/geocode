package main

import (
	"fmt"
	"time"
)

func greet(phrase string, done chan bool) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, done chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	close(done)
}

func main() {
	dones := make([]chan bool, 4)
	dones[0] = make(chan bool)
	go greet("Nice to meet you!", dones[0])
	dones[1] = make(chan bool)
	go greet("How are you?", dones[1])
	dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", dones[2])
	dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", dones[3])

	for _, done := range dones {
		<-done 
	}
}
