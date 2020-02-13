package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

func main() {
	var p1 person
	p1.age = 18
	p1.city = "甘肃"
	p1.name = "PFinal社区"
	fmt.Println(p1)
	p5 := person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Println(p5)
}
