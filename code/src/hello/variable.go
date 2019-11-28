package main
import (  
    "fmt"
    "math"
)


func main()  {
	a, b := 145.8, 543.8
    c := math.Min(a, b)
    fmt.Println("minimum value is ", c)
}