package main

import "fmt"

func main() {
	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)

	copy(a6, a4)
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	fmt.Println()

	b6 := []int{-10, 1, 2, 3, 4, 5}
	b4 := []int{-1, -2, -3, -4}
	fmt.Println("b6:", b6)
	fmt.Println("b4:", b4)
	copy(b4, b6)
	fmt.Println("b6:", b6)
	fmt.Println("b4:", b4)
	fmt.Println()
}
