package main

import "fmt"

func main() {
	experimentArrayAndSlice()
	experimentSharedArray()
	experimentAppend()
	experimentCopy()
	experimentPrediction()
	experimentSliceArgument()
}

func modifySlice(s []int) {
	s[0] = 99
	s = append(s, 4)
	fmt.Printf("inside: value=%v len=%d cap=%d\n", s, len(s), cap(s))
}

func experimentSliceArgument() {
	source := []int{1, 2, 3}

	modifySlice(source)

	fmt.Printf(
		"outside: value=%v len=%d cap=%d\n",
		source,
		len(source),
		cap(source),
	)
}

func experimentPrediction() {
	a := make([]int, 2, 3)
	a[0] = 1
	a[1] = 2
	fmt.Println("&a:", &a[1])

	b := a
	fmt.Println("&b:", &b[1])
	b = append(b, 3)
	b[0] = 100

	fmt.Println("a:", a)
	fmt.Println("a len:", len(a))
	fmt.Println("a cap:", cap(a))
	fmt.Println("&a:", &a[1])
	fmt.Println("&b:", &b[1])
	fmt.Println("b:", b)
	fmt.Println("b len:", len(b))
	fmt.Println("b cap:", cap(b))

	b = append(b, 4)
	b[1] = 200
	fmt.Println("&a0:", &a[0])
	fmt.Println("&b0:", &b[0])
	fmt.Println("&a1:", &a[1])
	fmt.Println("&b1:", &b[1])

	fmt.Println("a:", a)
	fmt.Println("b:", b)
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
	fmt.Printf("base=%v\n", base)
	fmt.Printf("base[:3]=%v\n", base[:3])

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
