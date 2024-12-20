package algorithms

import "errors"

type Node struct {
	val  interface{}
	prev *Node
	next *Node
}

func (n Node) Val() interface{} {
	return n.val
}

func (n Node) Prev() *Node {
	return n.prev
}

func (n Node) Next() *Node {
	return n.next
}

func NewNode(data interface{}) *Node {
	return &Node{
		val:  data,
		prev: nil,
		next: nil,
	}
}

type DoublyLinkedList struct {
	head *Node
	end  *Node
}

func (dll *DoublyLinkedList) Head() *Node {
	return dll.head
}

func (dll *DoublyLinkedList) End() *Node {
	return dll.end
}

func NewDoublyLinkedList() *DoublyLinkedList {
	dll := DoublyLinkedList{
		head: NewNode(0),
		end:  NewNode(0),
	}
	dll.head.next = dll.end
	dll.end.prev = dll.head

	return &dll
}

func (dll *DoublyLinkedList) Add(val interface{}) *Node {
	newNode := NewNode(val)
	temp := dll.head.next

	dll.head.next = newNode
	temp.prev = newNode
	newNode.next = temp

	return newNode
}

func (dll *DoublyLinkedList) Remove(node *Node) error {
	if node.prev == nil || node.next == nil {
		return errors.New("node dont have prev or next")
	}

	prevNode := node.prev
	nextNode := node.next

	prevNode.next = nextNode
	nextNode.prev = prevNode

	return nil
}
