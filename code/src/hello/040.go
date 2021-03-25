package main
import "fmt"

func main(){
	// 与数组不同  slice 的类型仅由它所含的元素决定
	// 要创建一个长度非零的空 sliec 需要使用内建的方法 make
	s := make([]string,3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	s = append(s,"d","e")
	fmt.Println(s)
	c :=make([]string,len(s))
	copy(c,s)
	fmt.Println(c)

	l :=s[2:3]
	fmt.Println(l)
	l = s[:2]
	fmt.Println(l)
	l = s[2:]
	fmt.Println(l)

	t :=[]string{"g","h","i"}
	fmt.Println(t)

}