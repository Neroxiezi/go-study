package main

import (
	"flag"
	"time"
)

func main() {
	// 定义命令行参数
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age,"age",18,"年龄")
	flag.BoolVar(&married, "married", false, "婚否")
}
