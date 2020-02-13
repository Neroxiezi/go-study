package main

import "fmt"
import "unsafe"

func main()  {
	a :=true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
    fmt.Println("c:", c)
    d := a || b
	fmt.Println("d:", d)
	
	var e int = 89
	f := 95
	fmt.Println("value of a is", e, "and b is", f)
	fmt.Printf("type of a is %T, size of a is %d\n", a, unsafe.Sizeof(e)) //type and size of a
	fmt.Printf("type of a is %T, size of a is %d\n", a, unsafe.Sizeof(f)) //type and size of a
}