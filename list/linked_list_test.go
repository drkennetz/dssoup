package list

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLinkedList_InsertToFront(t *testing.T) {
	a := createStarterList()
	a.Show()
	a.InsertDataToHead(4)
	a.Show()
	require.Equal(t, 4, a.head.data)
	require.Equal(t, 1, a.head.next.data)
	require.Equal(t, 2, a.head.next.next.data)
	require.Equal(t, 3, a.head.next.next.next.data)
}

func TestSList_InsertDataAfterFirstValue(t *testing.T) {
	a := createDupList()
	a.Show()
	err := a.InsertDataAfterFirstValue(2, 3)
	a.Show()
	require.Equal(t, 3, a.head.next.data)
	require.NoError(t, err)
	b := createDupList()
	err2 := b.InsertDataAfterFirstValue(1, 4)
	require.ErrorContains(t, err2, "value not found in list")
}

func createStarterList() *SList {
	a := NewSNode(1, nil)
	a.next = NewSNode(2, nil)
	a.next.next = NewSNode(3, nil)
	return NewLinkedList(a)
}

func createDupList() *SList {
	a := NewSNode(2, nil)
	a.next = NewSNode(2, nil)
	return NewLinkedList(a)
}
