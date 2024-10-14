package command

// CommandTwo is a concrete command to execute ActionTwo on the Receiver.
type CommandTwo struct {
	receiver *Receiver
}

// NewCommandTwo creates a new CommandTwo.
func NewCommandTwo(receiver *Receiver) *CommandTwo {
	return &CommandTwo{receiver: receiver}
}

// Execute performs the command's action.
func (c *CommandTwo) Execute() string {
	return c.receiver.ActionTwo()
}
