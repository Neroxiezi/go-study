package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {
	var emp4 Employee
	fmt.Println("Employee 4", emp4)
}
