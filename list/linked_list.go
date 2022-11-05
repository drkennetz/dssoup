package list

import (
	"fmt"
)

// LinkedList implements the Singly Linked list
type LinkedList struct {
	head *SNode
}

// NewLinkedList is a factory method for creating a new LinkedList
func NewLinkedList(h *SNode) *LinkedList {
	return &LinkedList{
		head: h,
	}
}

// Show Pretty prints a linked list
func (l *LinkedList) Show() {
	h := l.head
	if h != nil {
		fmt.Printf("%v ", h.data)
		h = h.next
	}
	for h != nil {
		fmt.Printf("-> %v ", h.data)
		h = h.next
	}
}

// InsertDataToHead creates a new node for data
// to be added to front of linked list
func (l *LinkedList) InsertDataToHead(newData interface{}) {
	newNode := NewSNode(newData, nil)
	newNode.next = l.head
	l.head = newNode
}

// InsertDataAfterFirstValue creates a new node for data
// and inserts node after first instance of value is found
func (l *LinkedList) InsertDataAfterFirstValue(v interface{}, newData interface{}) {
	
}
