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

	for {
		fmt.Println()
		fmt.Print("> ")
		scanner.Scan()
		input:= scanner.Text()
		inputArgs:= strings.Fields(input)
		command:= inputArgs[0]

		switch (command) {
		case "info":
			getInstruction()
		case "add":
			if len(inputArgs) != 5 {
				fmt.Println("Wrong count arguments for commands try again!")
				continue
			}

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
			id:= list.GetId()
			fmt.Println(id)
			list.Add(time.DateOnly, inputArgs[2], amount,id)
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
		case "summary":
			if len(inputArgs) == 1 {
				amount:= list.Summury()
				fmt.Println("amount of all expenses", amount)
			}

			if len(inputArgs) != 3 {
				continue
			} 

			if inputArgs[1] == "--month" {
					month, err := strconv.Atoi(inputArgs[2])

					if err != nil {
						fmt.Println("Something wrong please try again")
						continue
					}
					store.List.SummuryForMonth(month)
				}
			
		case "exit":
			fmt.Println()
			fmt.Println("We will hope you comeback later!")
			return
		default:
			getMessageWronComand()
			continue
		}
		
	}

}



