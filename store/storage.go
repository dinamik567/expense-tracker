package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Storage struct {
	filePath string
}

type Date struct {
	Expenses  ExpenseList
	Id int
}

type Store interface {
	NewStorage()
	ReadFile()
	SaveDate()
}

var store = NewStorage()

func NewStorage() *Storage {
    _, filename, _, _ := runtime.Caller(0)
    dir := filepath.Dir(filename)  // Путь к папке store/
    return &Storage{filepath.Join(dir, "store.json")}
}

func (storage *Storage) ReadFile() (Date, error) {
	file, err := os.ReadFile(storage.filePath)

	if err != nil {
		fmt.Println("Ошибка чтения, файла", err)
		return Date{}, err
	}

	var data Date

	err = json.Unmarshal(file, &data)

	if err != nil {
		fmt.Println("Error converting data to json format", err)
		os.Exit(1)
	}

	return data, err
}

func (storage *Storage) SaveDate(data Date) error {
	jsonObject, err:= json.Marshal(data)

	if err != nil {
		fmt.Println("Error converting data to json format")
		os.Exit(1)
	}

	err =  os.WriteFile(storage.filePath, jsonObject, 0644)

	if err != nil {
		fmt.Println("Ошибка записи файла", err)
		return err
	}

	return nil
}

