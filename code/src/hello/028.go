package main
import (
	"fmt"
)
// 创建结构体
type Saiyan struct {
	Name string
	Power int
}

// goku := Saiyan{
// 	Name string
// 	Power int
// }

func main() {
	goku := &Saiyan{"Power", 9000}
	//fmt.Println(goku)
	Super(goku)
	fmt.Println(goku.Power)
}

// func Super(s Saiyan) {
// 	s.Power += 10000
// }
func Super(s *Saiyan) {
	s.Power += 10000
}