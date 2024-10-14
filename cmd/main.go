package main

import (
	"command-pattern-go/pkg/command"
	"fmt"
)

func main() {
	receiver := &command.Receiver{}

	commandOne := command.NewCommandOne(receiver)
	commandTwo := command.NewCommandTwo(receiver)

	fmt.Println("Executing CommandOne:")
	fmt.Println(commandOne.Execute())

	fmt.Println("Executing CommandTwo:")
	fmt.Println(commandTwo.Execute())
}
