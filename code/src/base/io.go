package main

import (
	"io"
	"fmt"
	"strings"
	"os"
	"bufio"
)

func ReadForm(reader io.Reader,num int) ([]byte,error){
	p := make([]byte,num)
	n, err := reader.Read(p)
	if n > 0 {
        return p[:n], nil
    }
    return p, err
}

func main() {
	// filename := "text.txt";
	// dir, _ := os.Getwd()
	// fmt.Println(dir)
	// dirPath := filepath.Join(dir, filename)
	// fmt.Println(dirPath)
	// file,err := os.Open(filename)
	// if err != nil {
    //     fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
    //     return
    // }
	// data, err := ReadForm(file, 9)
	data, _ := ReadForm(strings.NewReader("from string"), 12)
	fmt.Println(data)

	reader := strings.NewReader("Go语言中文网")
	p := make([]byte, 6)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %d\n", p, n)

	// WriterAt

	file, err := os.Create("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Golang中文社区——这里是多余")
	n, err = file.WriteAt([]byte("Go语言中文网"), 24)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	//读取
	file, err = os.Open("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()

	//Seeker 接口
	reader = strings.NewReader("cccccccccc")
	reader.Seek(-6, io.SeekEnd)
	r, _, _ := reader.ReadRune()
	fmt.Printf("%c\n", r)
}