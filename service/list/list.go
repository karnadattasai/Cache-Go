/*Package list provides the data stucture list and supported methods for implementing cache*/
package list

import "fmt"

// Pair is the data type that list node holds
type Pair struct {
	Key   int
	Value int
}

// Node is the element in List data structure
type Node struct {
	P     Pair
	left  *Node
	right *Node
}

// List is the data structure that provides List properties
type List struct {
	front *Node
	back  *Node
	len   int
}

// Lister provides methods on list
type Lister interface {
	Remove(*Node)
	Push(P Pair) *Node
	MoveBack(*Node)
	Front() *Node
	Back() *Node
	Len() int
	Display()
}

// Display prints list
func (l *List) Display() {
	for i := l.front; i != nil; i = i.right {
		fmt.Printf("%d ", i.P.Key)
	}
	fmt.Printf("\n")
}

// MoveBack moves the given node to back
func (l *List) MoveBack(node *Node) {
	if node.right == nil {
		return
	}
	if node.left == nil {
		l.front = node.right
		l.front.left = nil
	} else {
		node.left.right = node.right
		node.right.left = node.left
	}
	node.left = l.back
	l.back.right = node
	l.back = node
	node.right = nil

}

// Push pushes an element in to the List at the back
func (l *List) Push(P Pair) *Node {
	node := Node{P, nil, nil}
	if l.front == nil && l.back == nil {
		l.front = &node
		l.back = l.front
		l.len = 1
		return &node
	}
	node.left = l.back
	l.back.right = &node
	l.back = &node
	l.len = l.len + 1
	return &node
}

// Remove deletes a given node
func (l *List) Remove(node *Node) {
	l.len = l.len - 1
	if node.left == nil && node.right == nil {
		l.front = nil
		l.back = nil
		return
	}
	if node.left == nil && node.right != nil {
		l.front = node.right
		l.front.left = nil
		return
	}
	if node.left != nil && node.right == nil {
		l.back = node.left
		l.back.right = nil
		return
	}
	node.left.right = node.right
	node.right.left = node.left
}

// Front return front node of the list
func (l *List) Front() *Node {
	return l.front
}

// Back return back node of the list
func (l *List) Back() *Node {
	return l.back
}

// Len return front node of the list
func (l *List) Len() int {
	return l.len
}
