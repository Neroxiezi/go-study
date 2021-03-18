package main
import (
	"fmt"
)

// 创建结构体
type Saiyan struct {
	Name string
	Power int
}

func main() {
	goku := &Saiyan {
		Name: "pfinal",
		Power:20,
	}
	fmt.Println(goku.Name)
}



//fmt.Println(goku.Name)