package main

import (
	"io"
	"fmt"
	"os"
)

func main() {
	
	file,err :=os.Open("./test.txt")
	if err !=nil {
		fmt.Println("open file failed!, err:", err)
		return 
	}
	defer file.Close()
	var tmp = make([]byte, 128)
	n,err := file.Read(tmp)
	
	if err !=nil {
		fmt.Println("读取文件错误")
		return	
	}	
	if err == io.EOF {
		fmt.Println("文件读取完毕")
		return	
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))
}