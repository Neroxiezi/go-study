package main

import (
	"io"
	"fmt"
	"os"
)

func main() {
	file,err :=os.Open("./test.txt")
	if err != nil {
		fmt.Println("读取文件错误")
		return 
	}
	defer file.Close()
	var content []byte
	var tmp = make([]byte, 128)
	for{
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)	
	}

	fmt.Println(string(content))
}