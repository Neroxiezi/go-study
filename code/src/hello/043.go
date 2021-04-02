package main

import "fmt"

func main() {
	// 多重赋值
	i := 10
	j := 20
	i, j = j, i
	fmt.Println(i,j)

	// 复数
	t := 3 + 5i
	fmt.Println(t)
	fmt.Printf("Type == %T \n", t)

	// 输入
	var tmp int
	fmt.Scanf("%d",&tmp)
	fmt.Println("tmp === ",tmp)

	var tmp2 int
	fmt.Scan(&tmp2)
	fmt.Printf("type == %T \n",tmp2)
}
