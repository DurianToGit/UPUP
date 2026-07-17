package main

import "fmt"

func main() {
	experimentArrayAndSlice()
	experimentSharedArray()
	experimentAppend()
	experimentCopy()
}

func experimentArrayAndSlice() {
	fmt.Println("=== array and slice ===")

	array := [3]int{10, 20, 30}
	slice := []int{10, 20, 30}

	fmt.Printf("array type=%T len=%d value=%v\n", array, len(array), array)
	fmt.Printf(
		"slice type=%T len=%d cap=%d value=%v\n",
		slice,
		len(slice),
		cap(slice),
		slice,
	)
}

func experimentSharedArray() {
	fmt.Println("\n=== shared underlying array ===")

	base := []int{10, 20, 30, 40, 50}
	left := base[1:4]
	right := base[2:5]

	fmt.Printf("before: base=%v left=%v right=%v\n", base, left, right)

	left[1] = 999

	fmt.Printf("after:  base=%v left=%v right=%v\n", base, left, right)
}

func experimentAppend() {
	fmt.Println("\n=== append and capacity ===")

	base := make([]int, 2, 4)
	base[0] = 10
	base[1] = 20

	child := base[:2]

	fmt.Printf(
		"before append: base=%v len=%d cap=%d child=%v len=%d cap=%d\n",
		base,
		len(base),
		cap(base),
		child,
		len(child),
		cap(child),
	)
	fmt.Printf("&base[0]=%p\n", &base[0])
	fmt.Printf("&child[0]=%p\n", &child[0])

	child = append(child, 30)

	fmt.Printf(
		"after first append: base=%v child=%v len=%d cap=%d\n",
		base,
		child,
		len(child),
		cap(child),
	)
	fmt.Printf("&base[0]=%p\n", &base[0])
	fmt.Printf("&child[0]=%p\n", &child[0])

	child = append(child, 40, 50)

	fmt.Printf(
		"after expansion: base=%v child=%v len=%d cap=%d\n",
		base,
		child,
		len(child),
		cap(child),
	)
	fmt.Printf("&base[0]=%p\n", &base[0])
	fmt.Printf("&child[0]=%p\n", &child[0])
}

func experimentCopy() {
	fmt.Println("\n=== independent copy ===")

	source := []int{10, 20, 30}

	target1 := source
	target2 := append([]int(nil), source...)

	target1[0] = 100
	target2[1] = 200

	fmt.Printf("source=%v\n", source)
	fmt.Printf("target1=%v\n", target1)
	fmt.Printf("target2=%v\n", target2)
}
