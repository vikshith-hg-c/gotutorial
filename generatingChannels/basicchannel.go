package main

import (
	"fmt"
	"math/rand"
	"time"
)

func process(ch chan int) {
	n := rand.Intn(5)
	time.Sleep(time.Duration(n) * time.Second)
	ch <- n
}

func main() {
	ch := make(chan int)
	go process(ch)

	fmt.Println("waiting for process")
	n := <-ch
	fmt.Printf("Process took %d seconds\n", n)
}
