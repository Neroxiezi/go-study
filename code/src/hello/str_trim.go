package main

import (  
    "fmt"
)

func noEmpty(data []string) []string {
	i := 0
	for _, v := range data {
		if v != "" {
			data[i] = v
			i++
		}
	}
	return data[:2]
}

func noReplace(data []string) []string  {
	tmp :=data[:1]
	outer:
	for _, v := range data {
		for i:=0; i<len(tmp);i++ {
			if v == tmp[i] {
				continue outer;
			}
		}
		tmp=append(tmp,v)
	}
	return tmp
}

func main() {
	data := []string{"red", "", "black", "red", "", "pink"}
	afterData := noEmpty(data)
	fmt.Println(afterData)
	afterDataOne := noReplace(data)
	fmt.Println(afterDataOne)
}
