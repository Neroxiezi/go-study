package main

import (
	"fmt"
	//"strconv"
)

// 基本数据类型
// Go 自爱不同类型的变量之间复制时需要显式转化

func main() {
	var i int32 = 100
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	fmt.Printf("i=%v n1=%v n2=%v n3=%v\n", i, n1, n2, n3)
	fmt.Printf("i type is %T\n",i)
	fmt.Printf("n1 type is %T\n",n1)

	// 基本数据类型相互转换
	// 1. Go 中 数据类型的转换可以实从表示范围小---> 表示范围大, 也可以范围大
	// 2. 被转换的是变量的数据  变量本身的数据类型并没有变化

	str := fmt.Sprintf("%d",i)
	fmt.Printf("str type is%T str %q\n",str,str)

	str2 := fmt.Sprintf("%f",n1)
	fmt.Printf("str type is%T str %q\n",str2,str2)

	isTrue := false
	fmt.Printf("str type %T str %v\n",isTrue,isTrue)
	str3 := fmt.Sprintf("%t",isTrue)
	fmt.Printf("str type %T str %q\n",str3,str3)

}
