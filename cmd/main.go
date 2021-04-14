package main

import (
	"os"
	"userservice/cmd/commands"

	_ "userservice/imports"
)

func main() {
	if err := commands.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
