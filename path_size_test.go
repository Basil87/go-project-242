package code

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSize(t *testing.T) {
	testdata := "testdata"

	t.Run("file", func(t *testing.T) {
		path := filepath.Join(testdata, "test.txt")

		size, err := GetSize(path, false, false)

		require.NoError(t, err)
		require.Equal(t, int64(17), size)
	})

	t.Run("empty file", func(t *testing.T) {
		path := filepath.Join(testdata, "empty.txt")

		size, err := GetSize(path, false, false)

		require.NoError(t, err)
		require.Equal(t, int64(0), size)
	})

	t.Run("directory", func(t *testing.T) {
		path := "testdata"

		size, err := GetSize(path, false, false)

		require.NoError(t, err)
		require.Equal(t, int64(34), size)
	})

	t.Run("directory recursive", func(t *testing.T) {
		path := "testdata"

		size, err := GetSize(path, false, true)

		require.NoError(t, err)
		require.Equal(t, int64(71), size)
	})

	t.Run("directory recursive with hidden files", func(t *testing.T) {
		path := "testdata"

		size, err := GetSize(path, true, true)

		require.NoError(t, err)
		require.Equal(t, int64(84), size)
	})

	t.Run("not exists", func(t *testing.T) {
		path := filepath.Join(testdata, "unknown")

		_, err := GetSize(path, false, false)

		require.Error(t, err)
	})
}
