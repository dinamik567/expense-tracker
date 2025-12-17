package main

import (
	"expense/traker/CLI"
	"expense/traker/store"
)



func main() {
	cli.Run(store.List)
}



