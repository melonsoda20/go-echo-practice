package files

import (
	"os"
)

func ReadFile(filePath string) (*os.File, error) {
	fileData, err := os.Open(filePath)
	return fileData, err
}
