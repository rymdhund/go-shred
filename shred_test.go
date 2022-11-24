package shred

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

// This file contains one test that makes sure that the helper function actually overwrite the contents of a file with the correct contents
// Other tests that would be nice but that I don't have time to implement:
// - Test that the file is removed
// - Test incorrect input like Shred on a dir or a non-existing file.
// - Full test of Shred() that would create and mount a filesystem that we could inspect afterwards to see that the contents of the file was overwritten.
// - Test that it works as expected on different file systems.
// - To actually know that the file was overwritten the correct number of times we could create a mock filesystem that recorded the operations.

func TestOverwrite(t *testing.T) {
	dir := t.TempDir()
	filePath := path.Join(dir, "file")
	f, err := os.Create(filePath)
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.Write([]byte("12345"))
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	data := strings.NewReader("abcdefghijklmnop")

	if err := overwriteFile(filePath, 2, data); err != nil {
		t.Fatal(err)
	}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "fghij" {
		t.Errorf("Exepected content to be overwritten, got \"%s\"", string(content))
	}
}
