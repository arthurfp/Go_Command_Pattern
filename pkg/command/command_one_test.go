package command

import "testing"

func TestCommandOne_Execute(t *testing.T) {
	receiver := &Receiver{}
	command := NewCommandOne(receiver)
	expected := "Receiver: Executing ActionOne"
	result := command.Execute()

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
