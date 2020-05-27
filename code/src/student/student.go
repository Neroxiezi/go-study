package main

import "fmt"

type student struct {
    id int
    name string
    class string
}

type studentMgr struct {
    allStudents []*student
}

func newStudent(id int,name, class string) *student {
    return &student{
        id:id,
        name:name,
        class:class,
    }
}
func newStudentMgr() *studentMgr {
    return &studentMgr{
        allStudents :make([]*student,0,100),
    }
}

func (s *studentMgr)addStudent(newStu *student) {
    s.allStudents = append(s.allStudents,newStu)
}


func (s *studentMgr) showStudent() {
    for _,v := range s.allStudents{
        fmt.Printf("学号:%d 姓名:%s 班级:%s \n" ,v.id,v.name,v.class)
    }
}
func (s *studentMgr) editStudent(newStu *student) {
     for i,v := range s.allStudents{
        if v.id == newStu.id {
            s.allStudents[i] = newStu
            return
        }
     }
    fmt.Println("没有要修改的学生")
}


