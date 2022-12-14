package slist

import (
	"errors"
	"fmt"
)

// SList implements the Singly Linked List
type SList[N comparable] struct {
	head  *sNode[N]
	count int
}

// NewLinkedList is a factory method for creating a new SList
func NewLinkedList[N comparable](h *sNode[N]) *SList[N] {
	if h == nil {
		return &SList[N]{
			head:  h,
			count: 0,
		}
	}
	return &SList[N]{
		head:  h,
		count: 1,
	}
}

// Show Pretty prints a linked list
func (l *SList[N]) Show() {
	h := l.head
	if h != nil {
		fmt.Printf("%v ", h.data)
		h = h.next
	} else {
		fmt.Printf("slist has no data")
	}
	for h != nil {
		fmt.Printf("-> %v ", h.data)
		h = h.next
	}
	fmt.Println("")
}

// InsertDataToHead creates a new node for data
// to be added to front of linked list
func (l *SList[N]) InsertDataToHead(newData N) {
	newNode := NewsNode(newData)
	newNode.next = l.head
	l.head = newNode
	l.count = l.count + 1
}

// InsertDataAfterFirstValue creates a new node for data
// and inserts node after first instance of value is found
func (l *SList[N]) InsertDataAfterFirstValue(v N, newData N) error {
	h := l.head
	newNode := NewsNode(newData)
	for h != nil {
		if h.data == v {
			newNode.next = h.next
			h.next = newNode
			l.count = l.count + 1
			return nil
		}
		h = h.next
	}
	return errors.New("value not found in list")
}

// Append adds a node to the end of a linked list
func (l *SList[N]) Append(newData N) {
	newNode := NewsNode(newData)
	if l.head == nil {
		l.head = newNode
		l.count = 1
		return
	}
	h := l.head
	for h.next != nil {
		h = h.next
	}
	h.next = newNode
	l.count = l.count + 1
}

// MergeSort MergeSorts a linked list
//  1. if the size of the linked list is 1 then return the head
//  2. Find mid using the tortoise and hare approach
//  3. Store the next of mid in head2 IE right sub-linked list
//  4. make the next midpoint null
//  5. recursively call MergeSort on both left and right sub
//     and store new head of left and right linked list
//  6. call merge() given the arguments new heads of left and right
//     sub lists and store the final head returned after merging
//  7. Return the final head of the merged linked list
func MergeSort() {}

// merge merges 2 linked lists
func merge() {}

// findMid finds the middle of a linked list
func findMid() {}
