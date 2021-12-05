package os

import (
	"os"
)

type FileManager interface {
	CreateDirectory(dirName string) error
	Exists(path string) (bool, error)
}

type fileManager struct {
}

func NewFileManager() FileManager {
	return fileManager{}
}

func (f fileManager) CreateDirectory(dirPath string) error {
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return err
	}

	return nil
}

func (f fileManager) Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
