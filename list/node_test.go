package list

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSNode_InsertNodeAfter(t *testing.T) {
	a := NewSNode(1, nil)
	a.InsertNodeAfter(2)
	require.Equal(t, a.next.data, 2)
}
