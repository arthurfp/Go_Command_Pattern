package command

// CommandOne is a concrete command to execute ActionOne on the Receiver.
type CommandOne struct {
	receiver *Receiver
}

// NewCommandOne creates a new CommandOne.
func NewCommandOne(receiver *Receiver) *CommandOne {
	return &CommandOne{receiver: receiver}
}

// Execute performs the command's action.
func (c *CommandOne) Execute() string {
	return c.receiver.ActionOne()
}
