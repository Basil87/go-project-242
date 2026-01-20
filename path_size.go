package code

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetSize(path string) (int64, error) {
	var total int64

	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
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

func HumanSize(size int64, human bool) string {
	const unit = 1000
	if !human || size < unit {
		return fmt.Sprintf("%dB", size)
	}

	div, exp := float64(unit), 0
	for n := float64(size) / unit; n >= unit && exp <= len(meterings) - 1; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f%s", float64(size)/div, meterings[exp])
}
