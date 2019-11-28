package main

import (
	"fmt"
)

func main() {
	// var a [3]int //int array with length 3
	// a[0] = 12    // array index starts at 0
	// a[1] = 78
	// a[2] = 50
	// fmt.Println(a)
	// b := [3]int{12, 78, 50} // short hand declaration to create array
	// fmt.Println(b)
	// c :=[3]int{12}
	// fmt.Println(c)
	// e:=[...]int{12,13,14,14}
	// fmt.Println(e)
	// a := [...]float64{67.7, 89.8, 21, 78}
	// sum := float64(0)
	// for i, v := range a { //range returns both the index and value
	// 	fmt.Printf("%d the element of a is %.2f\n", i, v)
	// 	sum += v
	// }
	// fmt.Println("\nsum of all elements of a", sum)
	// fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	// fruitslice := fruitarray[1:3]
	// fmt.Println(fruitslice,cap(fruitslice))
	i := make([]int, 5, 5)
	fmt.Println(i)
	cars := []string{"Ferrari", "Honda", "Ford"}
    fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
    cars = append(cars, "Toyota")
    fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))
}
