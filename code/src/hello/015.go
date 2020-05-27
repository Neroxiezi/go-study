package main
import "fmt"

func show(x interface{}) {
    fmt.Printf("type:%T value:%v\n", x, x)
}
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

func main(){
    s := "Hello PFinal"
    show(s)
    //var x interface{}
    //x ="Hello PFinal"
    justifyType(s)
}
