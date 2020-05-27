package main

import (
    "fmt"
)
type Employee struct {
    firstName, lastName string
    age, salary         int
}
func main(){
    emp3 := struct {
        firstName, lastName string
        age, salary         int
    }{
        firstName: "Andreah",
        lastName:  "Nikola",
        age:       31,
        salary:    5000,
    }
    fmt.Println("Employee 3", emp3)
    var emp4 Employee //zero valued structure
    fmt.Println("Employee 4", emp4)
    emp6 := Employee{"Sam", "Anderson", 55, 6000}
    fmt.Println("FirstName:",emp6.firstName)
    emp8 := &Employee{"Sam1", "Anderson", 55, 6000}
    fmt.Println("FirstName:",emp8.firstName)
    fmt.Println("FirstName:",(*emp8).firstName)
}
