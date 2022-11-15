package slist

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSList_InsertDataToHeadEmpty(t *testing.T) {
	a := NewLinkedList[int](nil)
	a.InsertDataToHead(4)
	require.Equal(t, a.head.data, 4)
	require.Equal(t, a.count, 1)
}

func TestSList_InsertDataToHead(t *testing.T) {
	a := createStarterList()
	require.Equal(t, a.count, 3)
	a.Show()
	a.InsertDataToHead(4)
	a.Show()
	require.Equal(t, a.count, 4)
	require.Equal(t, 4, a.head.data)
	require.Equal(t, 1, a.head.next.data)
	require.Equal(t, 2, a.head.next.next.data)
	require.Equal(t, 3, a.head.next.next.next.data)
}

func TestSList_InsertDataAfterFirstValueEmpty(t *testing.T) {
	a := NewLinkedList[int](nil)
	err := a.InsertDataAfterFirstValue(1, 2)
	require.ErrorContains(t, err, "value not found in list")
}

func TestSList_InsertDataAfterFirstValue(t *testing.T) {
	a := createDupList()
	a.Show()
	require.Equal(t, a.count, 2)
	err := a.InsertDataAfterFirstValue(2, 3)
	a.Show()
	require.Equal(t, a.count, 3)
	require.Equal(t, 3, a.head.next.data)
	require.NoError(t, err)
	b := createDupList()
	err2 := b.InsertDataAfterFirstValue(1, 4)
	require.ErrorContains(t, err2, "value not found in list")
}

func TestSList_Append(t *testing.T) {
	a := NewLinkedList[int](nil)
	require.Equal(t, a.count, 0)
	a.Show()
	a.Append(1)
	require.Equal(t, a.count, 1)
	require.Equal(t, a.head.data, 1)
	a.Show()
	a.Append(2)
	require.Equal(t, a.count, 2)
	require.Equal(t, a.head.next.data, 2)
	a.Show()
	a.Append(3)
	require.Equal(t, a.count, 3)
	require.Equal(t, a.head.next.next.data, 3)
	a.Show()
}

func createStarterList() *SList[int] {
	a := NewsNode(1)
	b := NewLinkedList(a)
	b.Append(2)
	b.Append(3)
	return b
}

func createDupList() *SList[int] {
	a := NewLinkedList[int](nil)
	a.Append(2)
	a.Append(2)
	return a
}
