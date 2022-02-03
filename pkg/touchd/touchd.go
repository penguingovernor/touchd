// package touchd implements a utility to create files and its necessary parent directories.
package touchd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const permDir = 0755

// CreateFiles calls CreateFile exactly once for each provided file.
// It returns nil if all files are created successfully.
// Otherwise, it returns a single reduced error.
func CreateFiles(fileNames ...string) error {
	// Keep track of errors.
	errStrBuilder := strings.Builder{}
	nErrors := 0

	// Iterate over all the files and create them, keeping track of errors.
	for _, file := range fileNames {
		if err := CreateFile(file); err != nil {
			nErrors++
			errStrBuilder.WriteString(fmt.Sprintf("[%s: %s] ", file, err))
		}
	}

	// If no errors, then yay!
	// Return early.
	if nErrors == 0 {
		return nil
	}

	// Otherwise, build the error message and return it.
	errStr := errStrBuilder.String()
	errStr = errStr[:len(errStr)-1] // Trims the last ' '

	// Handle the plural correctly.
	errStrPrefixBuilder := strings.Builder{}
	errStrPrefixBuilder.WriteString("failed to touch file")
	if nErrors > 1 {
		errStrPrefixBuilder.WriteRune('s')
	}

	return fmt.Errorf("%s: %s", errStrPrefixBuilder.String(), errStr)
}

// CreateFile creates a file and any required parent directories.
// If the file already exists, then the file's access/modified time is updated to reflect the current time.
//
// It returns nil, if the file already exists or is successfully created.
func CreateFile(fileName string) error {
	path := filepath.Dir(fileName)
	if err := os.MkdirAll(path, permDir); err != nil {
		return err
	}
	file, err := os.Open(fileName)
	if errors.Is(err, fs.ErrNotExist) {
		// File doesn't exist create it.
		createdFile, err := os.Create(fileName)
		if err != nil {
			return err
		}
		return createdFile.Close()
	}
	if err != nil {
		// Failed for some other reason...
		return err
	}
	defer file.Close()

	// File or directory exists...
	stats, err := file.Stat()
	if err != nil {
		return err
	}

	// File is actually directory,
	if stats.IsDir() {
		return errors.New("exists as directory")
	}

	// File exists, update modified/access time.
	t := time.Now().Local()
	return os.Chtimes(fileName, t, t)
}
