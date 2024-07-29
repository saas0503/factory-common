package futils

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func Test_ToUpper(t *testing.T) {
	t.Parallel()
	require.Equal(t, "/MY/NAME/IS/:PARAM/*", ToUpper("/my/name/is/:param/*"))
}

const (
	largeStr = "/RePos/GoFiBer/FibEr/iSsues/187643/CoMmEnts/RePos/GoFiBer/FibEr/iSsues/CoMmEnts"
	upperStr = "/REPOS/GOFIBER/FIBER/ISSUES/187643/COMMENTS/REPOS/GOFIBER/FIBER/ISSUES/COMMENTS"
	lowerStr = "/repos/gofiber/fiber/issues/187643/comments/repos/gofiber/fiber/issues/comments"
)

func Benchmark_ToUpper(b *testing.B) {
	var res string
	b.Run("futils", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			res = ToUpper(largeStr)
		}
		require.Equal(b, upperStr, res)
	})
	b.Run("IfToUpper-Upper", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			res = IfToUpper(upperStr)
		}
		require.Equal(b, upperStr, res)
	})
	b.Run("IfToUpper-Mixed", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			res = IfToUpper(largeStr)
		}
		require.Equal(b, upperStr, res)
	})
	b.Run("default", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			res = strings.ToUpper(largeStr)
		}
		require.Equal(b, upperStr, res)
	})
}

func Test_ToLower(t *testing.T) {
	t.Parallel()
	require.Equal(t, "/my/name/is/:param/*", ToLower("/MY/NAME/IS/:PARAM/*"))
	require.Equal(t, "/my1/name/is/:param/*", ToLower("/MY1/NAME/IS/:PARAM/*"))
	require.Equal(t, "/my2/name/is/:param/*", ToLower("/MY2/NAME/IS/:PARAM/*"))
	require.Equal(t, "/my3/name/is/:param/*", ToLower("/MY3/NAME/IS/:PARAM/*"))
	require.Equal(t, "/my4/name/is/:param/*", ToLower("/MY4/NAME/IS/:PARAM/*"))
}

func Benchmark_ToLower(b *testing.B) {
	var res string
	b.Run("futils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = ToLower(largeStr)
		}
		require.Equal(b, lowerStr, res)
	})
	b.Run("IfToLower-Lower", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = IfToLower(lowerStr)
		}
		require.Equal(b, lowerStr, res)
	})
	b.Run("IfToLower-Mixed", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = IfToLower(largeStr)
		}
		require.Equal(b, lowerStr, res)
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = strings.ToLower(largeStr)
		}
		require.Equal(b, lowerStr, res)
	})
}
func Test_IfToUpper(t *testing.T) {
	t.Parallel()
	require.Equal(t, "MYNAMEISPARAM", IfToUpper("MYNAMEISPARAM")) // already uppercase
	require.Equal(t, "MYNAMEISPARAM", IfToUpper("mynameisparam")) // lowercase to uppercase
	require.Equal(t, "MYNAMEISPARAM", IfToUpper("MyNameIsParam")) // mixed case
}

func Test_IfToLower(t *testing.T) {
	t.Parallel()
	require.Equal(t, "mynameisparam", IfToLower("mynameisparam"))           // already lowercase
	require.Equal(t, "mynameisparam", IfToLower("myNameIsParam"))           // mixed case
	require.Equal(t, "https://gofiber.io", IfToLower("https://gofiber.io")) // Origin Header Type URL
	require.Equal(t, "mynameisparam", IfToLower("MYNAMEISPARAM"))           // uppercase to lowercase
}
