package command

import "testing"

func TestReceiver_ActionOne(t *testing.T) {
	receiver := &Receiver{}
	expected := "Receiver: Executing ActionOne"
	result := receiver.ActionOne()

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestReceiver_ActionTwo(t *testing.T) {
	receiver := &Receiver{}
	expected := "Receiver: Executing ActionTwo"
	result := receiver.ActionTwo()

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
