package main

import "fmt"

func Multiply(a, b int, reply *int) {
	*reply = a * b
}


func main() {
	n := 0
	//reply := &n
	Multiply(10, 5, &n)
	fmt.Println("Multiply:", n) // Multiply: 50
}