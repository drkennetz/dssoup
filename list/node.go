package list

// SNode stores a key and the link to the next node
// SNode is used with singly linked lists
type SNode struct {
	data interface{}
	next *SNode
}

// NewSNode is a factory method for creating a new node
func NewSNode(data interface{}, next *SNode) *SNode {
	return &SNode{
		data: data,
		next: next,
	}
}

// InsertNodeAfter creates node from data
// and inserts after given node
func (n *SNode) InsertNodeAfter(data interface{}) {
	newNode := NewSNode(data, nil)
	newNode.next = n.next
	n.next = newNode
}
