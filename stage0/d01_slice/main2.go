package main

import "fmt"

func main() {
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
