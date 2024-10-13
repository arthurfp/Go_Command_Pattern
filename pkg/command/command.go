package command

// Command defines the interface for executing a command.
type Command interface {
	Execute() string
}
