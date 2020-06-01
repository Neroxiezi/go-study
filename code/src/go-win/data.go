package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func getWeiBoData() {
    var result string
	str, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
	file, err := os.Open(str + "/data/weibo.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	for {
		s, err := buf.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		result +=s
	}
	fmt.Println(result)
}

func main() {
	getWeiBoData()
}
