package shred

import (
	"crypto/rand"
	"errors"
	"io"
	"os"
)

// Shred will overwrite a given file with random data 3 times.
// The given path is expected to be a regular file.
func Shred(path string) error {
	if err := overwriteFile(path, 3, rand.Reader); err != nil {
		return err
	}
	return os.Remove(path)
}

// overwriteFile overwrites the content of file with random data.
// The contents are overwritten the given number of times.
func overwriteFile(path string, times int, data io.Reader) error {
	f, err := os.OpenFile(path, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return err
	}

	if !stat.Mode().IsRegular() {
		return errors.New("Can't shred non-regular file")
	}

	size := stat.Size()
	for i := 0; i < times; i++ {
		io.CopyN(f, data, size)
		if err = f.Sync(); err != nil {
			return err
		}
		if _, err := f.Seek(0, 0); err != nil {
			return err
		}
	}

	return nil
}
