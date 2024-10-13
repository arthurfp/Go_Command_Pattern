package command

import "fmt"

// Receiver defines an example receiver that performs operations.
type Receiver struct{}

// ActionOne performs the first operation.
func (r *Receiver) ActionOne() string {
	return "Receiver: Executing ActionOne"
}

// ActionTwo performs the second operation.
func (r *Receiver) ActionTwo() string {
	return "Receiver: Executing ActionTwo"
}

// Example of logging or other receiver methods
func (r *Receiver) Log(action string) string {
	return fmt.Sprintf("Receiver: Logging action %s", action)
}
