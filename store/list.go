package store

import (
	"fmt"
	"os"
	"strconv"
)

type Expense struct {
	Date        string `json:"date"`
	Description string `json:"description"`
	Amount      float64 `json:"amount"`
}

type ExpenseLister interface {
	Add(date string, description string, amount float64, id string)
	Delete(id string) bool
	Summury() float64
	GetId() string
	ShowList()
}

type ExpenseList map[string]Expense

func GetList() ExpenseList {
	var data, err = store.ReadFile()

	if err != nil {
		fmt.Println("Wrong reading file")
		os.Exit(1)
	}

	return data
}

var List = GetList()

func (expList ExpenseList) GetId() string {
	if len(expList) == 0 {
		return "1"
	} else {
		return strconv.FormatInt(int64(len(expList) + 1), 10)
	}
}

func (expList ExpenseList) Add(date string, description string, amount float64, id string) {
	expList[id] = Expense{date, description, amount}
	store.SaveDate(expList)
}

func (expList ExpenseList) ShowList() {
	fmt.Println()
	for ind, value:= range List {
		fmt.Println(ind, value.Date, value.Description, value.Amount)
	}
}

func (expList ExpenseList) Delete(id string) bool {
	_, ok := List[id]
	
	if ok {
		delete(expList, id)
		store.SaveDate(List)
	}
	
	return ok
}



func (expList ExpenseList) Summury() float64 {
	var sum float64 = 0
	for _, value := range expList {
		sum += value.Amount
	}

	return sum
}

// func (expList *ExpenseList) SummuryForMonth() {

// }