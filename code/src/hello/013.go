package main
import "fmt"

type Sayer interface {
    say() string
}

type Mover interface {
    mover() string
}

type dog struct {
	name string
}

func (d dog) say() string {
    return string(d.name) + ": 汪汪汪"
}
func (d dog) mover() string {
    return d.name+":跑起来"
}

func main(){
    var x Sayer
	var y Mover

	var a = dog{name: "旺财"}
	x = a
	y = a
	fmt.Println(x.say())
	fmt.Println(y.mover())
}
