package main

import (
	"fmt"
)

// func main() {
//     n := 5
//     for i := 0; i < n; i++ {
//         for j := 0; j <= i; j++ {
//             fmt.Print("*")
//         }
//         fmt.Println()
//     }
// }

/*func main() {
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break outer
			}
		}
	}
}*/
func main() {
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
}
