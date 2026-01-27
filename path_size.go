package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetPathSize returns the size of a file or directory.
//
// Flags behavior:
//   - recursive (-r): recursively calculate directory size
//   - all (-a): include hidden files
//   - human (-h): print size in human-readable format
//
// The result is returned as a formatted string.
// An error is returned if the path cannot be accessed.
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := getSize(path, all, recursive)
	if err != nil {
		return "", err
	}

	output := formatSize(size, human)
	return output, nil
}

func getSize(path string, all bool, recursive bool) (int64, error) {
	var total int64

	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !all && !info.IsDir() && strings.HasPrefix(info.Name(), ".") {
			return nil
		}

		if !recursive && info.IsDir() && p != path {
			return filepath.SkipDir
		}

		if !info.IsDir() {
			total += info.Size()
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return total, nil
}

var meterings = []string{"KB", "MB", "GB", "TB", "PB", "EB"}

func formatSize(size int64, human bool) string {
	const unit = 1000
	if !human || size < unit {
		return fmt.Sprintf("%dB", size)
	}

	div, exp := float64(unit), 0
	for n := float64(size) / unit; n >= unit && exp < len(meterings)-1; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f%s", float64(size)/div, meterings[exp])
}
