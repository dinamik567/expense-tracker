package store

import (
    "fmt"
    "os"
    "path/filepath"
    "runtime"
)

type Storage struct {
	filePath string
}

var MyStorage = Storage{"store.txt"}

func NewStorage() *Storage {
    _, filename, _, _ := runtime.Caller(0)
    dir := filepath.Dir(filename)  // Путь к папке store/
    return &Storage{filepath.Join(dir, "store.txt")}
}

func (storage *Storage) ReadFile() ([]byte, error) {
	file, err := os.ReadFile(storage.filePath)

	if err != nil {
		fmt.Println("Ошибка чтения, файла", err)
		return nil, err
	}

	return file, err
}

func (storage *Storage) SaveDate(data []byte) error {
	err:=  os.WriteFile(storage.filePath, data, 0644)
	if err != nil {
		fmt.Println("Ошибка записи файла", err)
		return err
	}

	return nil
}

