package main
import (  
    "fmt"
    "math"
)


func main()  {
	a, b := 145.8, 543.8
    c := math.Min(a, b)
	fmt.Println("minimum value is ", c)
	e :=25
	f :='a'
	sum :=e + int(f)
	fmt.Println(sum)
	i := 10
    var j float64 = float64(i) //this statement will not work without explicit conversion
    fmt.Println("j", j)

}