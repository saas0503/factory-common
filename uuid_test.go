package futils

import (
	"crypto/rand"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_UUID(t *testing.T) {
	t.Parallel()
	res := UUID()
	require.Len(t, res, 36)
	require.NotEqual(t, "00000000-0000-0000-0000-000000000000", res)
}

func Benchmark_UUID(b *testing.B) {
	var res string
	b.Run("futils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = UUID()
		}
		require.Len(b, res, 36)
	})
	b.Run("default", func(b *testing.B) {
		rnd := make([]byte, 16)
		_, _ = rand.Read(rnd)
		for n := 0; n < b.N; n++ {
			res = fmt.Sprintf("%x-%x-%x-%x-%x", rnd[0:4], rnd[4:6], rnd[6:8], rnd[8:10], rnd[10:])
		}
		require.Len(b, res, 36)
	})
}
