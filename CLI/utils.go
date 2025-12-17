package cli

import (
	"strings"
	"fmt"
)

type Flags struct {
	description string
	amount string
}

var flags = Flags{"--description", "--amount"}


func getInstruction() {
	fmt.Println(`
add       Add a new expense (flags: --description, --amount, --date)
list      List expenses (flags: --month, --year)
summary   Show totals (flags: --month, --year)
delete    Delete expense by ID (flags: --id)
info    Show app info and stats
exit    Exit the application`)
}

func getMessageWronComand() {
	fmt.Println("Wrong command")
	fmt.Println("Typing 'info' for help")
}

func getMessageWronFlag(flag string) {
	fmt.Println("Wrong command invalid flag: ", flag)
}

func getMessageWronValidInputValue(typeValue string) {
	fmt.Println("Wrong valid input value: ", typeValue)
}

func checkValidInputValueForEmpty(input string) bool {
	s:= strings.TrimSpace(input)

	return len(s) > 0 
}