package main

import (
	"command-pattern-go/pkg/command"
	"fmt"
)

func main() {
	receiver := &command.Receiver{}

	fmt.Println(receiver.ActionOne())
	fmt.Println(receiver.ActionTwo())
}
