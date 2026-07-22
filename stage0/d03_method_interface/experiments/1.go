package experiments

import "fmt"

type User struct {
	ID   int
	Name string
}

func (u User) RenameByValue(name string) {
	u.Name = name
	fmt.Println("inside value method:", u)
}

func (u *User) RenameByPointer(name string) {
	u.Name = name
	fmt.Println("inside pointer method:", u)
}

// 实验一：值接收者与指针接收者
func ExperimentReceiver() {
	fmt.Println("=== receiver ===")

	user := User{
		ID:   1,
		Name: "Tom",
	}

	user.RenameByValue("Value Jerry")
	fmt.Println("after value method:", user)

	user.RenameByPointer("Pointer Jerry")
	fmt.Println("after pointer method:", user)
}
