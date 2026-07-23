package experiments

import "testing"

func TestDoSomethingReturnsTypedNil(t *testing.T) {
	err := doSomething(true)

	if err == nil {
		t.Fatal("expected typed nil interface to be non-nil")
	}
}

func TestDoSomethingFixedReturnsNil(t *testing.T) {
	err := doSomethingFixed(true)

	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}
