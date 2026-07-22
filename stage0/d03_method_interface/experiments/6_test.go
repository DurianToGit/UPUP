package experiments

import "testing"

type FakeNotifier struct {
	Messages []string
}

func (f *FakeNotifier) Notify(message string) error {
	f.Messages = append(f.Messages, message)
	return nil
}

func TestCompleteOrder(t *testing.T) {
	notifier := &FakeNotifier{}
	service := NewOrderService(notifier)

	err := service.CompleteOrder(1001)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(notifier.Messages) != 1 {
		t.Fatalf(
			"expected 1 message, got %d",
			len(notifier.Messages),
		)
	}

	if notifier.Messages[0] != "order 1001 completed" {
		t.Fatalf(
			"unexpected message: %s",
			notifier.Messages[0],
		)
	}
}
