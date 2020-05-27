package main
import "fmt"

type Mover interface {
    move() string
}

type dog struct {}
func (d dog) move() string {
    return "狗狗移动嘞"
}

func main(){
    var x Mover
    var wangcai = dog{}
    x = wangcai
    fmt.Println(x.move())
    var fugui = &dog{}
    x = fugui
    fmt.Println(x.move())
}

