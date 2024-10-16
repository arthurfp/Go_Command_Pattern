package command

import "testing"

func TestInvoker_ExecuteAll(t *testing.T) {
	receiver := &Receiver{}
	commandOne := NewCommandOne(receiver)
	commandTwo := NewCommandTwo(receiver)

	invoker := &Invoker{}
	invoker.AddCommand(commandOne)
	invoker.AddCommand(commandTwo)

	results := invoker.ExecuteAll()
	expected := []string{
		"Receiver: Executing ActionOne",
		"Receiver: Executing ActionTwo",
	}

	for i, result := range results {
		if result != expected[i] {
			t.Errorf("Expected %s but got %s", expected[i], result)
		}
	}
}
