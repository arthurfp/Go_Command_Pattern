package main

import (
	"command-pattern-go/pkg/command"
	"fmt"
)

func main() {
	receiver := &command.Receiver{}

	commandOne := command.NewCommandOne(receiver)
	commandTwo := command.NewCommandTwo(receiver)

	invoker := &command.Invoker{}
	invoker.AddCommand(commandOne)
	invoker.AddCommand(commandTwo)

	fmt.Println("Executing all commands via Invoker:")
	for _, result := range invoker.ExecuteAll() {
		fmt.Println(result)
	}
}
