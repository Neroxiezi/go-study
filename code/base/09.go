package main

import "fmt"

func main() {
	integer := make([]int, 20)
	integer = append(integer, -5900)
	for i := 0; i < len(integer); i++ {
		fmt.Println(i)
	}
	s2 := integer[1:3]
	s2[0] = -100
	s2[1] = 123456
	fmt.Println(s2)
}
