package experiments

import "fmt"

func inspect(value any) {
	fmt.Printf(
		"value=%v type=%T nil=%v\n",
		value,
		value,
		value == nil,
	)
}

func ExperimentInterfaceValue() {
	fmt.Println("\n=== interface value ===")

	var empty any
	inspect(empty)

	var number int = 10
	inspect(number)

	var userPointer *User
	inspect(userPointer)

	var userPointer2 *User = nil
	var value any = userPointer2

	// 直接打印
	fmt.Println(userPointer2 == nil)
	fmt.Println(value == nil)
	// 使用 inspect
	inspect(userPointer2)
	inspect(value)
}
