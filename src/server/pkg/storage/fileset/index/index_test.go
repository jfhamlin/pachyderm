package index

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/pachyderm/pachyderm/src/client/pkg/require"
	"github.com/pachyderm/pachyderm/src/server/pkg/obj"
	"github.com/pachyderm/pachyderm/src/server/pkg/storage/chunk"
	"github.com/pachyderm/pachyderm/src/server/pkg/storage/fileset/tar"
)

const (
	testPath = "test"
)

func write(tb testing.TB, objC obj.Client, chunks *chunk.Storage, fileNames []string) {
	iw := NewWriter(context.Background(), objC, chunks, testPath)
	for _, fileName := range fileNames {
		hdr := &Header{
			Hdr: &tar.Header{Name: fileName},
			Idx: &Index{
				DataOp: &DataOp{},
			},
		}
		require.NoError(tb, iw.WriteHeaders([]*Header{hdr}))
	}
	require.NoError(tb, iw.Close())
}

func actualFiles(tb testing.TB, objC obj.Client, chunks *chunk.Storage, opts ...Option) []string {
	ir := NewReader(context.Background(), objC, chunks, testPath, opts...)
	defer func() {
		require.NoError(tb, ir.Close())
	}()
	result := []string{}
	for {
		hdr, err := ir.Next()
		if err == io.EOF {
			return result
		}
		require.NoError(tb, err)
		result = append(result, hdr.Hdr.Name)
	}
}

func expectedFiles(fileNames []string, prefix string) []string {
	result := []string{}
	for _, fileName := range fileNames {
		if strings.HasPrefix(fileName, prefix) {
			result = append(result, fileName)
		}
	}
	return result
}

func pathRange(fileNames []string) *PathRange {
	return &PathRange{
		Lower: fileNames[0],
		Upper: fileNames[len(fileNames)-1],
	}
}

func Check(t *testing.T, permString string) {
	objC, chunks := chunk.LocalStorage(t)
	defer func() {
		chunk.Cleanup(objC, chunks)
		objC.Delete(context.Background(), testPath)
	}()
	fileNames := Generate(permString)
	averageBits = 12
	write(t, objC, chunks, fileNames)
	t.Run("Full", func(t *testing.T) {
		expected := fileNames
		actual := actualFiles(t, objC, chunks)
		require.Equal(t, expected, actual)
	})
	t.Run("FirstFile", func(t *testing.T) {
		prefix := fileNames[0]
		expected := []string{prefix}
		actual := actualFiles(t, objC, chunks, WithPrefix(prefix))
		require.Equal(t, expected, actual)
		actual = actualFiles(t, objC, chunks, WithRange(pathRange(expected)))
		require.Equal(t, expected, actual)
	})
	t.Run("FirstRange", func(t *testing.T) {
		prefix := string(fileNames[0][0])
		expected := expectedFiles(fileNames, prefix)
		actual := actualFiles(t, objC, chunks, WithPrefix(prefix))
		require.Equal(t, expected, actual)
		actual = actualFiles(t, objC, chunks, WithRange(pathRange(expected)))
		require.Equal(t, expected, actual)
	})
	t.Run("MiddleFile", func(t *testing.T) {
		prefix := fileNames[len(fileNames)/2]
		expected := []string{prefix}
		actual := actualFiles(t, objC, chunks, WithPrefix(prefix))
		require.Equal(t, expected, actual)
		actual = actualFiles(t, objC, chunks, WithRange(pathRange(expected)))
		require.Equal(t, expected, actual)
	})
	t.Run("MiddleRange", func(t *testing.T) {
		prefix := string(fileNames[len(fileNames)/2][0])
		expected := expectedFiles(fileNames, prefix)
		actual := actualFiles(t, objC, chunks, WithPrefix(prefix))
		require.Equal(t, expected, actual)
		actual = actualFiles(t, objC, chunks, WithRange(pathRange(expected)))
		require.Equal(t, expected, actual)
	})
	t.Run("LastFile", func(t *testing.T) {
		prefix := fileNames[len(fileNames)-1]
		expected := []string{prefix}
		actual := actualFiles(t, objC, chunks, WithPrefix(prefix))
		require.Equal(t, expected, actual)
		actual = actualFiles(t, objC, chunks, WithRange(pathRange(expected)))
		require.Equal(t, expected, actual)
	})
	t.Run("LastRange", func(t *testing.T) {
		prefix := string(fileNames[len(fileNames)-1][0])
		expected := expectedFiles(fileNames, prefix)
		actual := actualFiles(t, objC, chunks, WithPrefix(prefix))
		require.Equal(t, expected, actual)
		actual = actualFiles(t, objC, chunks, WithRange(pathRange(expected)))
		require.Equal(t, expected, actual)
	})
}

func TestSingleLevel(t *testing.T) {
	Check(t, "abc")
}

func TestMultiLevel(t *testing.T) {
	Check(t, "abcdefg")
}
