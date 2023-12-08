package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleep() {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
}
func producer(ch chan<- int) {
	for {
		sleep()
		n := rand.Intn(100)
		fmt.Println(n)
		ch <- n
	}
}

func consumer(ch <-chan int, name string) {
	for n := range ch {
		fmt.Printf(" %s <- %d\n", name, n)
	}
}

func fanout(chA <-chan int, chB, chC chan<- int) {
	for n := range chA {
		if n < 50 {
			chB <- n
		} else {
			chC <- n
		}
	}
}

func main() {
	chA := make(chan int, 6)
	chB := make(chan int, 6)
	chC := make(chan int, 6)

	go producer(chA)
	go consumer(chB, "B")
	go consumer(chC, "C")

	fanout(chA, chB, chC)
}
