package main

type Queue struct {
	items chan int
}

func (q *Queue) Enqueue(i int) {
	q.items <- i
}

func (q *Queue) Dequeue() int {

	return <-q.items
}

func main() {
	q := Queue{
		items: make(chan int, 4),
	}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	println(q.Dequeue())
	println(q.Dequeue())
	println(q.Dequeue())
	println(q.Dequeue())

}
