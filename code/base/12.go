package main

import "fmt"

func main() {
	fmt.Println()
	array4 := [4]int{4, -4, 4, -4}
	s6 := []int{1, -1, 1, -1, -5, 5}
	copy(s6, array4[0:])
	fmt.Println("array4:", array4[0:])
	fmt.Printf("s6:", s6)
	fmt.Println()
	array5 := [5]int{5, -5, 5, -5, 5}
	s7 := []int{7, 7, -7, 7, -7, 7}
	copy(array5[0:], s7)
	fmt.Println("array5:", array5)
	fmt.Printf("s7:", s7)
	fmt.Println()
}
