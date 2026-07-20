package main

import (
	"fmt"
	"sync"
)

func main() {
	experimentMapAssignment()
	experimentMapLookup()
	experimentNilMap()
	experimentStruct()
	experimentMapStruct()
	experimentConcurrentMapWrite()
}

// Map 的共享行为
func experimentMapAssignment() {
	fmt.Println("=== map assignment ===")

	nilMap := map[string]int{}
	fmt.Println("nilMap:", nilMap)

	original := map[string]int{
		"apple":  10,
		"banana": 20,
	}

	copied := original
	copied["apple"] = 99
	copied["orange"] = 30

	fmt.Println("original:", original)
	fmt.Println("copied:", copied)
}

// 零值与存在性
func experimentMapLookup() {
	fmt.Println("\n=== map lookup ===")

	scores := map[string]int{
		"Tom": 0,
	}

	tomScore := scores["Tom"]
	jerryScore := scores["Jerry"]

	fmt.Println("Tom:", tomScore)
	fmt.Println("Jerry:", jerryScore)

	tomScore, tomExists := scores["Tom"]
	jerryScore, jerryExists := scores["Jerry"]

	fmt.Printf("Tom: value=%d exists=%v\n", tomScore, tomExists)
	fmt.Printf("Jerry: value=%d exists=%v\n", jerryScore, jerryExists)
}

// 实验三：nil Map
func experimentNilMap() {
	fmt.Println("\n=== nil map ===")

	var data map[string]int

	fmt.Println("data is nil:", data == nil)
	fmt.Println("read missing key:", data["missing"])

	// data["new"] = 1
}

// 实验四：Struct 值与指针
type User struct {
	ID   int
	Name string
}

func modifyUserValue(user User) {
	user.Name = "Value Changed"
}

func modifyUserPointer(user *User) {
	user.Name = "Pointer Changed"
}

func experimentStruct() {
	fmt.Println("\n=== struct value and pointer ===")

	user := User{
		ID:   1,
		Name: "Original",
	}

	modifyUserValue(user)
	fmt.Println("after value:", user)

	modifyUserPointer(&user)
	fmt.Println("after pointer:", user)
}

// 实验五：Map 中的 Struct
func experimentMapStruct() {
	fmt.Println("\n=== map struct ===")
	// 原始代码
	/*users := map[int]User{
		1: {
			ID:   1,
			Name: "Tom",
		},
	}*/
	/*
		users[1].Name = "Jerry"
	*/
	// 方案一 取出、修改、写回
	// user := users[1]
	// user.Name = "Jerry"
	// users[1] = user
	// 方案二 Map 存储指针
	users := map[int]*User{
		1: {
			ID:   1,
			Name: "Tom",
		},
	}

	users[1].Name = "Jerry"
	fmt.Println("users:", users[1])
}

// 实验六：Map 并发写

type SafeMap struct {
	mu   sync.RWMutex
	data map[int]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[int]int),
	}
}
func (m *SafeMap) Set(key, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}
func (m *SafeMap) Get(key int) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, exists := m.data[key]
	return value, exists
}
func (m *SafeMap) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.data)
}

func experimentConcurrentMapWrite() {
	data := NewSafeMap()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				data.Set(workerID*1000+j, j)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("map size:", data.Len())
}
