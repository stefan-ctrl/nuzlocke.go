package util

import "os"

func ReadFile(filePath string) (*[]byte, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func WriteFile(path string, bytes []byte) error {
	return os.WriteFile(path, bytes, 0644)
}
