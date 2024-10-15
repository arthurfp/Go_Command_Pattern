package command

// Invoker stores and executes commands.
type Invoker struct {
	commands []Command
}

// AddCommand adds a command to the invoker's list.
func (i *Invoker) AddCommand(cmd Command) {
	i.commands = append(i.commands, cmd)
}

// ExecuteAll executes all stored commands in order.
func (i *Invoker) ExecuteAll() []string {
	var results []string
	for _, cmd := range i.commands {
		results = append(results, cmd.Execute())
	}
	return results
}
