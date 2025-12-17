package store

import "strconv"

type Expense struct {
	date        string
	description string
	amount      float64
}

type ExpenseLister interface {
	Add(date string, description string, amount float64, id string)
	Delete(id string)
	Summury() float64
	GetId() string
}

type ExpenseList map[string]Expense

var list = make(ExpenseList)

func GetList() ExpenseList {
	return list
}

func (expList ExpenseList) GetId() string {
	if len(expList) == 0 {
		return "1"
	} else {
		return strconv.FormatInt(int64(len(expList) + 1), 10)
	}
}

func (expList ExpenseList) Add(date string, description string, amount float64, id string) {
	expList[id] = Expense{date, description, amount}
}

func (expList ExpenseList) Delete(id string) {
	delete(expList, id)
}

func (expList ExpenseList) Summury() float64 {
	var sum float64 = 0
	for _, value := range expList {
		sum += value.amount
	}

	return sum
}

// func (expList *ExpenseList) SummuryForMonth() {

// }

// ??? не уверен что метод необходим
// func CreateNewExp(date string, description string, amount float64) Expense {
// 	exp := Expense{date, description, amount}

// 	return exp
// }