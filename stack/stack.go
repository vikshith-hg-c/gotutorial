package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	left := len(s.items)
	if left == 0 {
		return -1
	}
	item, items := s.items[left-1], s.items[0:left-1]
	s.items = items
	return item
}

func main() {
	a := Stack{}
	a.Push(1)
	a.Push(2)
	a.Push(3)
	a.Push(4)
	a.Push(5)
	a.Push(6)

	fmt.Println(a.Pop())
	fmt.Println(a.Pop())
	fmt.Println(a.Pop())
	fmt.Println(a.Pop())
	fmt.Println(a.Pop())
	fmt.Println(a.Pop())
	fmt.Println(a.Pop())
}
