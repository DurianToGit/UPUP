package experiments

import "fmt"

type Notifier interface {
	Notify(message string) error
}

type EmailNotifier struct{}

func (n EmailNotifier) Notify(message string) error {
	fmt.Println("email:", message)
	return nil
}

type SMSNotifier struct{}

func (n SMSNotifier) Notify(message string) error {
	fmt.Println("sms:", message)
	return nil
}

type OrderService struct {
	notifier Notifier
}

func NewOrderService(notifier Notifier) *OrderService {
	return &OrderService{
		notifier: notifier,
	}
}

func (s *OrderService) CompleteOrder(orderID int) error {
	message := fmt.Sprintf("order %d completed", orderID)
	return s.notifier.Notify(message)
}

func InjectNotifier() {
	// notifier := EmailNotifier{}
	notifier := SMSNotifier{}
	order := NewOrderService(notifier)
	err := order.CompleteOrder(1)
	if err != nil {
		fmt.Println("error:", err)
	}
}
