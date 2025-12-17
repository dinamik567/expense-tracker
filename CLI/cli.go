package cli

import (
	"bufio"
	"expense/traker/store"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type CLI struct {}

func Run(list store.ExpenseLister) {
	
	scanner:= bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Expensive Tracker")
	fmt.Println("Typing 'info' for help")
	fmt.Print("> ")
	
	for {
		scanner.Scan()
		input:= scanner.Text()
		inputArgs:= strings.Fields(input)
		command:= inputArgs[0]

		switch (command) {
		case "info":
			getInstruction()
		case "add":
			if (inputArgs[1] != flags.description) {
				getMessageWronFlag(flags.description)
				continue
			}

			if (inputArgs[3] != flags.amount) {
				getMessageWronFlag(flags.amount)
				continue
			}

			if (!checkValidInputValueForEmpty(inputArgs[2])) {
				getMessageWronValidInputValue("description")
				continue
			}

			if (!checkValidInputValueForEmpty(inputArgs[4])) {
				getMessageWronValidInputValue("amount")
				continue
			}

			amount, err:= strconv.ParseFloat(inputArgs[4], 64)

			if err != nil {
				getMessageWronValidInputValue("amount")
				continue
			}

			list.Add(time.DateOnly, inputArgs[2], amount, list.GetId())
			continue
		case "list":
			list.ShowList()
		case "delete":
			if (len(inputArgs) != 2) {
				getMessageWronValidInputValue("id")
				continue
			}
			if (!checkValidInputValueForEmpty(inputArgs[1])) {
				getMessageWronValidInputValue("id")
				continue
			}

			ok:= list.Delete(inputArgs[1])

			if !ok {
				fmt.Println("Something wrong try again")
			}
		case "exit":
			fmt.Println()
			fmt.Println("We will hope you comeback later!")
			return
		default:
			getMessageWronComand()
			continue
		}
		
		fmt.Println()
		fmt.Print("> ")
	}

}



