package slist

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewsNode(t *testing.T) {
	a := NewsNode(1)
	require.Equal(t, 1, a.data)
}
