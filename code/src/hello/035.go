package main

import "fmt"

func min(s ...int) int {
	if len(s)==0 {
		return 0
	}

	min :=s[0]
	for _,v :=range s {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	x := min(1,2,2,0)
	fmt.Printf("The mininum is: %d\n", x)
	slice :=[]int{7,9,3,5,1}
	fmt.Println(slice)
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d \n", x)
}