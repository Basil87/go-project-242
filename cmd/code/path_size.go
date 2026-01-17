package code

import (
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
