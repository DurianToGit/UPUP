package main

import (
	"errors"
	"fmt"
)

var ErrInsufficientBalance = errors.New("insufficient balance")

func Withdraw(balance, amount int64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid withdraw amount: %d", amount)
	}

	if amount > balance {
		return fmt.Errorf(
			"withdraw amount=%d balance=%d: %w",
			amount,
			balance,
			ErrInsufficientBalance,
		)
	}

	return nil
}

func WithdrawW2V(balance, amount int64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid withdraw amount: %d", amount)
	}

	if amount > balance {
		return fmt.Errorf(
			"withdraw amount=%d balance=%d: %v",
			amount,
			balance,
			ErrInsufficientBalance,
		)
	}

	return nil
}

// ErrorIs
func experimentErrorIs() {
	err := Withdraw(100, 200)

	fmt.Println("error:", err)
	fmt.Println(
		"is insufficient balance:",
		errors.Is(err, ErrInsufficientBalance),
	)
}

// ErrorAs

type FieldError struct {
	Field   string
	Value   string
	Message string
}

func (e *FieldError) Error() string {
	return fmt.Sprintf(
		"field=%s value=%q: %s",
		e.Field,
		e.Value,
		e.Message,
	)
}

func ValidateUsername(username string) error {
	if username == "" {
		return &FieldError{
			Field:   "username",
			Value:   username,
			Message: "cannot be empty",
		}
	}

	return nil
}

func CreateUser(username string) error {
	if err := ValidateUsername(username); err != nil {
		return fmt.Errorf("create user failed: %w", err)
	}

	return nil
}

func experimentErrorAs() {
	err := CreateUser("")

	var fieldErr *FieldError

	fmt.Println("error:", err)
	if errors.As(err, &fieldErr) {
		fmt.Printf(
			"field=%s message=%s\n",
			fieldErr.Field,
			fieldErr.Message,
		)
	}
}
