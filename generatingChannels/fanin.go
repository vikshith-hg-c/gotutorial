package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleep() {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
}
func producer(ch chan<- int, name string) {
	for {
		sleep()
		n := rand.Intn(100)
		fmt.Printf("channel %s -> %d\n", name, n)
		ch <- n
	}
}

func consumer(ch <-chan int) {
	for n := range ch {
		fmt.Printf("c <- %d\n", n)
	}
}

func fanin(chA, chB <-chan int, chC chan<- int) {
	for chA != nil || chB != nil {
		select {
		case n, ok := <-chA:
			if ok {
				chC <- n
			} else {
				chA = nil
			}
		case n, ok := <-chB:
			if ok {
				chC <- n
			} else {
				chB = nil
			}
		}
	}
	close(chC)
}
func main() {
	chA := make(chan int, 6)
	chB := make(chan int, 6)
	chC := make(chan int, 6)

	go producer(chA, "AAAA")
	go producer(chB, "BBBB")
	go consumer(chC)
	fanin(chA, chB, chC)

}
