package main
import "fmt"

type Cat struct{}

func (c Cat) Say() string {
    return "喵喵喵"
}
type Dog struct{}
func (d Dog) Say() string {
    return "汪汪汪"
}

type Sayer interface {
    say()
}
func main(){
    c := Cat{}
    fmt.Println(c.Say())
    d := Dog{}
    fmt.Println(d.Say())
}
