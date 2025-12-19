package store

import (
	"fmt"
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

func GetDate() (ExpenseList, int) {
	var data, _ = store.ReadFile()

	return data.Expenses, data.Id
}

var List, NextId = GetDate()

func (expList ExpenseList) GetId() string {
	if len(expList) == 0 {
		return "1"
	} else {
		return strconv.FormatInt(int64(len(expList) + 1), 10)
	}
}


func (expList ExpenseList) Add(date string, description string, amount float64, id string) {
	List[id] = Expense{date, description, amount}
	store.SaveDate(Date{List, NextId + 1})
}

func (expList ExpenseList) ShowList() {
	fmt.Println()
	for ind, value:= range List {
		fmt.Println(ind, value.Date, value.Description, value.Amount)
	}
}

func (expList ExpenseList) Delete(id string) bool {
	_, ok := expList[id]
	
	if ok {
		delete(expList, id)
		store.SaveDate(Date{List, NextId})
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

func (expList *ExpenseList) SummuryForMonth(month int) {
	date, err:= getTime(int64(month))

	if err != nil {
		fmt.Println("Error format an input time")
	}

	var sum float64 = 0
	

	for _, expense:= range List {
		expenseDate, _:= getTimeFromString(expense.Date)
		if (expenseDate.Month() == date.Month()) {
			sum+= expense.Amount
		}
	}

	fmt.Println("Expense for input month",sum)
}