package main

import (
	"fmt"
	"time"
)

func main() {
	// 当前时间
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n", now)

	// 获取年月日
	fmt.Printf("年=%+v\n", now.Year())
	fmt.Printf("月=%+v\n", now.Month())
	fmt.Printf("月=%+v\n", int(now.Month()))
	fmt.Printf("日%+v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	// 格式化日期时间
	fmt.Printf("当前年月日 %d-%d", fmt)

}