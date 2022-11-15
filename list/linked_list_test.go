package list

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSList_InsertDataToHeadEmpty(t *testing.T) {
	a := NewLinkedList(nil)
	a.InsertDataToHead(4)
	require.Equal(t, a.head.data, 4)
}

func TestSList_InsertDataToHead(t *testing.T) {
	a := createStarterList()
	a.Show()
	a.InsertDataToHead(4)
	a.Show()
	require.Equal(t, 4, a.head.data)
	require.Equal(t, 1, a.head.next.data)
	require.Equal(t, 2, a.head.next.next.data)
	require.Equal(t, 3, a.head.next.next.next.data)
}

func TestSList_InsertDataAfterFirstValueEmpty(t *testing.T) {
	a := NewLinkedList(nil)
	err := a.InsertDataAfterFirstValue(1, 2)
	require.ErrorContains(t, err, "value not found in list")
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

func TestSList_Append(t *testing.T) {
	a := NewLinkedList(nil)
	a.Show()
	a.Append(1)
	require.Equal(t, a.head.data, 1)
	a.Show()
	a.Append(2)
	require.Equal(t, a.head.next.data, 2)
	a.Show()
	a.Append(3)
	require.Equal(t, a.head.next.next.data, 3)
	a.Show()
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
