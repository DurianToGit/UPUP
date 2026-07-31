package main

import "fmt"

func experimentDeferEvaluation() {
	fmt.Println("=== defer evaluation ===")

	value := 1

	defer fmt.Println("defer argument:", value)

	defer func() {
		fmt.Println("defer closure:", value)
	}()

	value = 2

	fmt.Println("normal:", value)
}

func CalculatePrice(price int) (result int) {
	result = price

	defer func() {
		result += 10
	}()

	return result
}
func CalculatePrice2(price int) (result int) {
	result = price

	defer func(result int) {
		result += 10
	}(result)

	return result
}

func ExperimentNamedReturn() {
	fmt.Println(CalculatePrice(100))
}

// 循环中的defer
type Resource struct {
	ID int
}

func OpenResource(id int) *Resource {
	fmt.Println("open:", id)
	return &Resource{ID: id}
}

func (r *Resource) Close() {
	fmt.Println("close:", r.ID)
}

func processResourcesBad(ids []int) {
	for _, id := range ids {
		resource := OpenResource(id)
		defer resource.Close()

		fmt.Println("process:", id)
	}
}

func processOneResource(id int) {
	resource := OpenResource(id)
	defer resource.Close()

	fmt.Println("process:", id)
}

func processResourcesGood(ids []int) {
	for _, id := range ids {
		processOneResource(id)
	}
}
