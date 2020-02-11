package main

import (
	"fmt"
)

func main() {
	// num := 10
	// if num % 2 == 0 { //checks if number is even
	//     fmt.Println("the number is even")
	// }  else {
	//     fmt.Println("the number is odd")
	// }
	// if num := 10; num % 2 == 0 { //checks if number is even
	//     fmt.Println(num,"is even")
	// }  else {
	//     fmt.Println(num,"is odd")
	// }
	//上面的代码中num 在if 语句中初始化。需要注意的一件事是，num只能从if和else内部访问。也就是说，num 的范围仅限于if else代码块。如果我们试图从if 或else外部访问num ，编译器就会报错。

	num := 99
	if num <= 50 {
		fmt.Println("number is less than or equal to 50")
	} else if num >= 51 && num <= 100 {
		fmt.Println("number is between 51 and 100")
	} else {
		fmt.Println("number is greater than 100")
	}
}
