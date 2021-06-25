package main

import (
	"fmt"
)

func main() {
	// 数组的遍历
	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	for k, v := range nums {
		fmt.Println(k, v, " ")
	}
	fmt.Println()

	// 切片的遍历
	slis := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for k, v := range slis {
		fmt.Println(k, v, " ")
	}
	fmt.Println()

	// 字典的遍历
	tmpMap := map[string]string{
		"a": "小明",
		"b": "小红",
		"c": "小张",
	}
	for k, v := range tmpMap {
		// k 为键值，v 为对应值
		fmt.Println(k, v, " ")
	}
}
