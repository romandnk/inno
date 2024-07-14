package testdata

import (
	"os"
	"strconv"
)

func CreateTempFiles(dir string, num int) ([]*os.File, error) {
	files := make([]*os.File, 0, num)
	for range num {
		f, err := os.CreateTemp(dir, strconv.Itoa(num))
		if err != nil {
			return nil, err
		}
		files = append(files, f)
	}
	return files, nil
}
