package main
import (
    "fmt"
    "reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}
func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}
func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}
    t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	printMethod(stu1)



}
