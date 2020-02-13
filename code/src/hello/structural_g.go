package main

import "fmt"

type person struct {
	name string
	age  int8
}

func newPerson(name string, age int8) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func (p person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func main() {
	p1 := newPerson("小王子", 22)
	fmt.Println(p1)
	p1.Dream()
}
