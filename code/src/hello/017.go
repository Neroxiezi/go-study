package main
import (
    "fmt"
    "reflect"
)

func main(){
    b :=1
    // *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	a = &b
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
}
