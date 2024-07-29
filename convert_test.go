package futils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnsafeString(t *testing.T) {
	t.Parallel()
	res := UnsafeString([]byte("Hello, World!"))
	require.Equal(t, "Hello, World!", res)
}
