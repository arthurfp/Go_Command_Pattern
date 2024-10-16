package command

import "testing"

func TestCommandTwo_Execute(t *testing.T) {
	receiver := &Receiver{}
	command := NewCommandTwo(receiver)
	expected := "Receiver: Executing ActionTwo"
	result := command.Execute()

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
