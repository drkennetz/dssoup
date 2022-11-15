package slist

// sNode stores a key and the link to the next node
// sNode is used with singly linked lists
// sNode is not exported because they should not be created by
// anything except a linked list
// this way, it's harder to create rogue linked lists
type sNode[N comparable] struct {
	data N
	next *sNode[N]
}

// NewsNode is a factory method for creating a new node
func NewsNode[N comparable](data N) *sNode[N] {
	return &sNode[N]{
		data: data,
		next: nil,
	}
}
