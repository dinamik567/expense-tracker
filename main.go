package main

import (
	"expense/traker/CLI"
	"expense/traker/store"
)



func main() {
	
	list:= store.GetList()
	cli.Run(list)

}



