package main

type Tree struct {
	node *Node
}

func (t *Tree) insert(value int) *Tree {
	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		t.node.insert(value)
	}
	return t
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) insert(value int) {

	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.insert(value)
		}
	}
}

func (n *Node) exists(value int) bool {
	if n == nil {
		return false
	}
	if n.value == value {
		return true
	}
	if value <= n.value {
		return n.left.exists(value)
	} else {
		return n.right.exists(value)
	}

}

func printNode(n *Node) {
	if n == nil {
		return
	}
	println(n.value)
	printNode(n.left)
	printNode(n.right)
}

func main() {
	t := &Tree{}
	t.insert(10).
		insert(8).
		insert(11).
		insert(13).
		insert(9).
		insert(25).
		insert(30).
		insert(5)

	printNode(t.node)

	println(t.node.exists(10))
}
