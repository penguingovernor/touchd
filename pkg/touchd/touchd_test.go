// package touchd implements a utility to create files and its necessary parent directories.
package touchd

import (
	"os"
	"strings"
	"testing"
)

func createFileName(basename string, dir ...string) string {
	builder := strings.Builder{}
	for _, parent := range dir {
		builder.WriteString(parent)
		builder.WriteRune(os.PathSeparator)
	}
	builder.WriteString(basename)
	return builder.String()
}

func TestCreateFile(t *testing.T) {
	t.Run("create regular file", func(t *testing.T) {
		fileName := createFileName("foo", os.TempDir())
		t.Cleanup(func() {
			os.RemoveAll(fileName)
		})
		if err := CreateFile(fileName); err != nil {
			t.Fatalf("got %s, wanted nil", err)
		}
		f, err := os.Open(fileName)
		if err != nil {
			t.Fatalf("got %s, wanted nil", err)
		}
		f.Close()
	})
	t.Run("create nested file", func(t *testing.T) {
		fileName := createFileName("foo", os.TempDir(), "bar")
		t.Cleanup(func() {
			os.RemoveAll(createFileName("bar", os.TempDir()))
		})
		if err := CreateFile(fileName); err != nil {
			t.Fatalf("got %s, wanted nil", err)
		}
		f, err := os.Open(fileName)
		if err != nil {
			t.Fatalf("got %s, wanted nil", err)
		}
		f.Close()
	})
}

func TestCreateFiles(t *testing.T) {
	t.Run("create multiple nested files", func(t *testing.T) {
		fileNameA := createFileName("bar", os.TempDir(), "foo")
		fileNameB := createFileName("baz", os.TempDir(), "foo")
		t.Cleanup(func() {
			os.RemoveAll(createFileName("foo", os.TempDir()))
		})
		if err := CreateFiles(fileNameA, fileNameB); err != nil {
			t.Fatalf("got %s, wanted err", err)
		}
		// FilenameA should exist
		f, err := os.Open(fileNameA)
		if err != nil {
			t.Fatalf("got %s, wanted nil", err)
		}
		f.Close()
	})

	t.Run("create file with conflicting inner dir", func(t *testing.T) {
		fileNameA := createFileName("bar", os.TempDir(), "foo")
		fileNameB := createFileName("baz", os.TempDir(), "foo", "bar")
		t.Cleanup(func() {
			os.RemoveAll(createFileName("foo", os.TempDir()))
		})
		if err := CreateFiles(fileNameA, fileNameB); err == nil {
			t.Fatalf("got nil, wanted err")
		}
		// FilenameA should exist
		f, err := os.Open(fileNameA)
		if err != nil {
			t.Fatalf("got %s, wanted nil", err)
		}
		f.Close()
	})

	t.Run("create file with conflicting outer dir", func(t *testing.T) {
		fileNameA := createFileName("bar", os.TempDir(), "foo")
		fileNameB := createFileName("foo", os.TempDir())
		t.Cleanup(func() {
			os.RemoveAll(createFileName("foo", os.TempDir()))
		})
		if err := CreateFiles(fileNameA, fileNameB); err == nil {
			t.Fatalf("got nil, wanted err")
		}
		// FilenameA should exist
		f, err := os.Open(fileNameA)
		if err != nil {
			t.Fatalf("got %s, wanted nil", err)
		}
		f.Close()
	})

	t.Run("create already existing file", func(t *testing.T) {
		fileNameA := createFileName("test", os.TempDir())
		t.Cleanup(func() {
			os.RemoveAll(fileNameA)
		})
		if err := CreateFiles(fileNameA, fileNameA); err != nil {
			t.Fatalf("got %s, wanted nil", err)
		}
		// FilenameA should exist
		f, err := os.Open(fileNameA)
		if err != nil {
			t.Fatalf("got %s, wanted nil", err)
		}
		f.Close()
	})

	t.Run("multiple errors", func(t *testing.T) {
		fileNameA := createFileName("bar", os.TempDir(), "foo")
		fileNameB := createFileName("baz", os.TempDir(), "foo", "bar")

		fileNameC := createFileName("2", os.TempDir(), "1")
		fileNameD := createFileName("3", os.TempDir(), "1", "2")
		t.Cleanup(func() {
			os.RemoveAll(createFileName("foo", os.TempDir()))
			os.RemoveAll(createFileName("1", os.TempDir()))
		})
		if err := CreateFiles(fileNameA, fileNameB, fileNameC, fileNameD); err == nil {
			t.Fatalf("got nil, wanted err")
		}
		// FilenameA should exist
		f, err := os.Open(fileNameA)
		if err != nil {
			t.Fatalf("got %s, wanted nil", err)
		}
		f.Close()
		// FilenameC should exist
		g, err := os.Open(fileNameC)
		if err != nil {
			t.Fatalf("got %s, wanted nil", err)
		}
		g.Close()
	})
}
