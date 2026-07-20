package main

import (
	"fmt"
	"sync"
	"testing"
)

// 实验三：nil Map的写入panic测试
func TestWriteNilMapShouldPanic(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic, but no panic occurred")
		} else {
			t.Log("panic occurred")
		}
	}()

	var data map[string]int
	data["new"] = 1
}

// 1. 测试不存在的Key返回 exists=false；
// 2. 测试nil Map写入会panic；
// 3. 测试SafeMap并发写入后长度为10000。
func TestSafeMap_Get(t *testing.T) {
	m := NewSafeMap()
	m.Set(1, 1)
	value, exists := m.Get(1)
	if !exists {
		t.Fatal("expected exists=true, but got exists=false")
	}
	if value != 1 {
		t.Fatalf("expected value=1, but got value=%d", value)
	}
	if m.Len() != 1 {
		t.Fatalf("expected len=1, but got len=%d", m.Len())
	}
	value2, exists2 := m.Get(2)
	if exists2 {
		t.Fatal("expected exists=false, but got exists=true")
	}
	if value2 != 0 {
		t.Fatalf("expected value=0, but got value=%d", value2)
	}

}

func TestNewSafeMapLength(t *testing.T) {
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
	if data.Len() != 10000 {
		t.Fatalf("expected len=10000, but got len=%d", data.Len())
	}

}
