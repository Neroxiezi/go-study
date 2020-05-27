package main
import (
    "fmt"
    "reflect"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n,kind:%v\n", v,v.Kind())
}
func main(){
    s := "Hello go"
    reflectType(s)
    var a float32 = 3.14
    reflectType(a)
    b :=true
    reflectType(b)
    var c int64 = 12
    reflectType(c)
    var d *float32
    reflectType(d)
    var e rune
    reflectType(e)
}
