package main

import "fmt"

func main() {
	ids := []int{33,76,75,63,43,23}
	for i, id := range ids {
		fmt.Println(i, "-", id)
	}

	// pointer
	a := 5
	b := &a
	fmt.Println(a, b)

	// memory address
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// Use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	*b = 10
	fmt.Println(a)
}
