package main

import "fmt"

func main() {
	// 要创建一个 空 map 需要使用 make

	m := make(map[string]int)
	// 使用典型的 `make[key] = val` 语法来设置键值对。
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println(m)

	// 使用 key 来打印map 的建值对
	fmt.Println(m["k1"])

	// 当从一个map 中取值得时候 可选的第二返回值指示这个建
	_, prs :=m["k2"]
	fmt.Println(prs)
}
