package main

import "fmt"

func calculateBill(price int, no int) int {  
    var totalPrice = price * no
    return totalPrice
}


func rectProps(length, width float64) (float64,float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}


func main()  {
	sum :=calculateBill(1,2)
	fmt.Println(sum)	
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f\n", area, perimeter) 
	area1, _ := rectProps(10.8, 5.6)
	fmt.Printf("Area %f\n ", area1)
}

