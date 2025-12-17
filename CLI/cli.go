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
	list.ShowList()
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
				getInstruction()
				continue
			}

			if (inputArgs[3] != flags.amount) {
				getMessageWronFlag(flags.amount)
				getInstruction()
				continue
			}

			if (!checkValidInputValueForEmpty(inputArgs[2])) {
				getMessageWronValidInputValue("description")
				getInstruction()
				continue
			}

			if (!checkValidInputValueForEmpty(inputArgs[4])) {
				getMessageWronValidInputValue("amount")
				getInstruction()
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
			
		case "exit":
			fmt.Println()
			fmt.Println("We will hope you comeback later!")
			return
		default:
			getMessageWronComand()
			continue
		}
		
		fmt.Print("> ")
	}

}



