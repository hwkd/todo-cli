package main

import (
	"fmt"
	"os"
	"todocli/internal/args"
)

func main() {
	result, err := args.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}

	switch result.Action {
	case args.ActionAdd:
		fmt.Println("You tried to add")
	case args.ActionUpdate:
		fmt.Println("You tried to update")
	case args.ActionDelete:
		fmt.Println("You tried to delete")
	case args.ActionMarkComplete:
		fmt.Println("You tried to mark complete")
	case args.ActionMarkIncomplete:
		fmt.Println("You tried to mark incomplete")
	}
}
