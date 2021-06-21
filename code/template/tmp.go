package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
}

func main() {
	s1, _ := template.ParseFiles("header.tmpl","content.tmpl","footer.tmpl")
	//s1.ExecuteTemplate(os.Stdout, "header", nil)
	//fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", nil)
}
