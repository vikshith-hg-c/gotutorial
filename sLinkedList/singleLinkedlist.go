package singlelinkedlist

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}

type Node struct {
	value int
	next  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func main() {
	a := &List{}
	a.Push(1)
	a.Push(2)

	b := &List{}
	b.Push(100)
	b.Push(101)

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

}
