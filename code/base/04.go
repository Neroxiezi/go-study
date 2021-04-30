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
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d\n", now.Year(),int(now.Month()),now.Day(),now.Hour(),now.Minute(),now.Second())
    // 2 使用 time.Format() 方法完成
	fmt.Printf(now.Format("2021/4/30 10:54:38"))
	fmt.Println()
	// 这个字符串的各个数字是固定的  必须是这样写
	// 这个字符串各个数字可以自由的组合
	fmt.Printf(now.Format("15|04|05"))
	fmt.Println()

	ｉ := 0
	for {
		i++
		fmt.Println(i)
		// 休眠
		time.Sleep(time.Millisecond * 100)
		if i == 100 {
			break
		}
	}

 

}