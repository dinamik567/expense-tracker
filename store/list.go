package store

type Expense struct {
	date        string
	description string
	amount      float64
}

type ExpenseLister interface  {
	Add(date string, description string, amount float64, id string)
	Delete(id string)
	Summury()float64
}

type ExpenseList map[string]Expense

var id = 0

func GetId() int {
	id++
	return id
}

var list = make(ExpenseList)

func GetList() ExpenseList {
	return list
}

func (expList ExpenseList) Add(date string, description string, amount float64, id string) {
	expList[id] = Expense{date ,description, amount}
}

func (expList ExpenseList) Delete(id string) {
	delete(expList, id)
}

func (expList ExpenseList) Summury()float64 {
	var sum float64 = 0
	for _, value:= range expList {
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