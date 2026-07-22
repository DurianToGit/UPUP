package experiments

import "fmt"

func printValue(value any) {
	switch v := value.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	case *User:
		if v == nil {
			fmt.Println("*User: nil")
			return
		}

		fmt.Println("*User:", v.Name)
	default:
		fmt.Printf("unknown type: %T\n", value)
	}
}

func ExperimentTypeAssertion() {
	values := []any{
		100,
		"hello",
		&User{Name: "Tom"},
		(*User)(nil),
		3.14,
	}

	for _, value := range values {
		printValue(value)
	}

	// 另外的类型断言测试
	value := any("hello")

	// 断言失败
	// text := value.(int)
	// fmt.Println(text)

	// 断言成功
	text := value.(string)
	fmt.Println(text)

	number, ok := value.(int)
	fmt.Println(number, ok)
}
