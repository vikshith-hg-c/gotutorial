package main

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func main() {
	a := &List{}
	a.Push(1)
	a.Push(2)
	a.Push(3)

	b := &List{}
	b.Push(100)
	b.Push(101)
	b.Push(102)

	n := a.First()
	m := b.First()
	for {
		println(n.value)
		println(m.value)
		n = n.Next()
		m = m.Next()
		if m == nil || n == nil {
			break
		}

	}

	x := a.Last()
	y := b.Last()
	for {

		println(x.value)
		println(y.value)
		x = x.Prev()
		y = y.Prev()
		if n == nil || y == nil {
			break
		}

	}

}
