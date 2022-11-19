package slist

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
//     This is a safe implementation which uses nlogn time, and logn additional space
func MergeSort[N comparable](l *SList[N]) *sNode[N] {
	head := l.head
	if head.next == nil {
		return head
	}
	mid := findMid(head)
}

// merge merges 2 linked lists
func merge[N comparable](head1, head2 *sNode[N]) *sNode[N] {
	merged := &sNode[N]{}
	tmp := merged
	for (head1 != nil) && (head2 != nil) {
		if head1.data < head2.data {
			tmp.next = head1
			head1 = head1.next
		} else {
			tmp.next = head2
			head2 = head2.next
		}
		tmp = tmp.next
	}
}

// findMid finds the middle of a linked list
func findMid[N comparable](head *sNode[N]) *sNode[N] {
	slow := head
	fast := head.next
	for (fast != nil) && (fast.next != nil) {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
