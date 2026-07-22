package experiments

import "fmt"

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func doSomething(success bool) error {
	var customErr *CustomError

	if !success {
		customErr = &CustomError{
			Message: "operation failed",
		}
	}

	return customErr
}

func doSomethingFixed(success bool) error {
	if success {
		return nil
	}

	return &CustomError{
		Message: "operation failed",
	}
}

func ExperimentTypedNil() {
	// err := doSomething(true)
	err := doSomethingFixed(true)

	fmt.Printf("err=%v type=%T nil=%v\n", err, err, err == nil)
}
