// producer -->channel --> worker pool --> channel -->
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//func reciver(ch <-chan int) {

//}

func worker(in, out chan int, id int) {
	for {
		n, ok := <-in
		if ok {
			time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
			fmt.Printf("%d processing in %d worker\n", n, id)
			out <- n
		} else {
			fmt.Println("closing out")
			close(out)
		}
	}
}

func producer(ch chan<- int) {
	i := 0
	for {
		fmt.Printf("sent %d\n", i)
		ch <- i
		i++
	}
	//close(ch)
}

func main() {
	in := make(chan int)
	out := make(chan int)

	go producer(in)
	//go reciver(out)
	for i := 0; i < 3; i++ {
		go worker(in, out, i)
	}
	for n := range out {
		fmt.Printf("Recived %d\n", n)
	}

}
