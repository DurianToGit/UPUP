package main

import (
	"errors"
	"fmt"
	"testing"
)

// errors.Is可以识别包装后的ErrInsufficientBalance；
func TestErrorIs(t *testing.T) {
	err := Withdraw(100, 200)
	if errors.Is(err, ErrInsufficientBalance) == false {
		t.Fatalf("expected error is insufficient balance, but got error: %v", err)
	}
}

// 将%w改成%v后，errors.Is无法识别原错误；
func TestErrorIsW2V(t *testing.T) {
	err := WithdrawW2V(100, 200)
	if errors.Is(err, ErrInsufficientBalance) == true {
		t.Fatalf("改成V之后不应该有insufficient balance的错误")
	}
}

// errors.As可以提取FieldError；
func TestErrorAs(t *testing.T) {
	err := CreateUser("")
	if err != nil {
		var fieldErr *FieldError
		if errors.As(err, &fieldErr) {
			fmt.Printf("field=%s message=%s\n", fieldErr.Field, fieldErr.Message)
		}
	}
}

// defer按照后进先出顺序执行

func TestDefer(t *testing.T) {
	experimentDeferEvaluation()
}

// defer参数求值和闭包读取变量的结果不同
func TestCalculatePrice(t *testing.T) {
	value := 100
	result1 := CalculatePrice(value)
	result2 := CalculatePrice2(value)
	t.Logf("result1: %v, result2: %v", result1, result2)
	if result1 != 110 {
		t.Fatalf("结果应该为110，这里为%v，发生了错误", result1)
	}
	if result1 == result2 {
		t.Fatalf("参数求值和闭包读取变量的结果应该不同，这里却相同，发生了错误")
	}
}
