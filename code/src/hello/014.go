package main
import "fmt"

// Mover 接口
type Mover interface {
	move()
}
type dog struct {
	name string
}
type car struct {
	brand string
}

func (d dog) move() {
    fmt.Printf("%s会跑\n", d.name)
}

func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}

func main(){
	var x Mover
	var a = dog{name: "旺财"}
	var b = car{brand: "保时捷"}
	x = a
	x.move()
	x = b
	x.move()
}
