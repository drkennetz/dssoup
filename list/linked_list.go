package list

import (
	"errors"
	"fmt"
)

// SList implements the Singly Linked list
type SList struct {
	head *SNode
}

// NewLinkedList is a factory method for creating a new SList
func NewLinkedList(h *SNode) *SList {
	return &SList{
		head: h,
	}
}

// Show Pretty prints a linked list
func (l *SList) Show() {
	h := l.head
	if h != nil {
		fmt.Printf("%v ", h.data)
		h = h.next
	} else {
		fmt.Printf("list has no data")
	}
	for h != nil {
		fmt.Printf("-> %v ", h.data)
		h = h.next
	}
	fmt.Println("")
}

// InsertDataToHead creates a new node for data
// to be added to front of linked list
func (l *SList) InsertDataToHead(newData interface{}) {
	newNode := NewSNode(newData, nil)
	newNode.next = l.head
	l.head = newNode
}

// InsertDataAfterFirstValue creates a new node for data
// and inserts node after first instance of value is found
func (l *SList) InsertDataAfterFirstValue(v interface{}, newData interface{}) error {
	h := l.head
	newNode := NewSNode(newData, nil)
	for h != nil {
		if h.data == v {
			newNode.next = h.next
			h.next = newNode
			return nil
		}
		h = h.next
	}
	return errors.New("value not found in list")
}

// Append adds a node to the end of a linked list
func (l *SList) Append(newData interface{}) {
	newNode := NewSNode(newData, nil)
	if l.head == nil {
		l.head = newNode
		return
	}
	h := l.head
	for h.next != nil {
		h = h.next
	}
	h.next = newNode
}
