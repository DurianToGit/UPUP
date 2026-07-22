package experiments

import "testing"

func TestDoSomething(t *testing.T) {
	err := doSomething(true)
	if err == nil {
		t.Fatalf("expected error, but got nil")
	}
}
