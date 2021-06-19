package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))

	var str1 string = "Hi, I'm PFinal, Hi."

	fmt.Printf("The position of \"PFinal\" is: ")
	fmt.Printf("%d\n", strings.Index(str1, "PFinal"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str1, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str1, "Hi"))

	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str1, "Burger"))

	var str2  string = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"
	fmt.Printf("Number of H's in %s is: ", str2)
	fmt.Printf("%d\n", strings.Count(str2, "H"))

	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))
}