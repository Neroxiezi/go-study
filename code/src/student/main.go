package main
import (
    "fmt"
    "os"
)
func showMenu(){
    fmt.Println("学生管理系统:")
    fmt.Println("1,学生列表")
    fmt.Println("2,添加学生")
    fmt.Println("3,编辑学生")
    fmt.Println("4,退出")
}
func getInput() *student{
   var (
   		id    int
   		name  string
   		class string
   	)
   	fmt.Println("请按要求输入学员信息")
   	fmt.Print("请输入学员的学号：")
   	fmt.Scanf("%d\n", &id)
   	fmt.Print("请输入学员的姓名：")
   	fmt.Scanf("%s\n", &name)
   	fmt.Print("请输入学员的班级：")
   	fmt.Scanf("%s\n", &class)
   	// 就能拿到用户输入的学员的所有信息
   	stu := newStudent(id, name, class) // 调用student的构造函数造一个学生
   	return stu

}
func main(){
    sm := newStudentMgr()
    for{
        //展示菜单
        showMenu()
        var input int
        fmt.Print("请输入要操作的序号:")
        fmt.Scanf("%d\n",&input)
        switch input {
        //学生列表
        case 1:
            sm.showStudent()
        // 添加学生
        case 2:
            stu := getInput()
            sm.addStudent(stu)
            sm.showStudent()
        // 编辑学生
        case 3:
            stu :=getInput()
            sm.editStudent(stu)
            sm.showStudent()
        // 退出
        case 4:
            os.Exit(0)
        }
    }

}
