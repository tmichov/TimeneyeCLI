package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tmichov/TimeneyeCLI/cmd"
)

func main() {
	args := os.Args[1:]

	err := cmd.SetupCommands()
	if err != nil {
		panic(err)
	}

	if len(args) == 0 {
		fmt.Println("Please provide a command")
		return
	}

	command, ok := cmd.Commands[args[0]]
	if !ok {
		fmt.Println("Command not found")
		fmt.Println("Try 'help' to get a list of available commands")
		return
	}

	err = command.Executor(args[0], args[1:])
	if err != nil {
		log.Fatal(err)
	}

	return
}
