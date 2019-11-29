package main

import (  
    "fmt"
)

func main() {  
    // personSalary := make(map[string]int)
    // personSalary["steve"] = 12000
    // personSalary["jamie"] = 15000
	// personSalary["mike"] = 9000
	personSalary := map[string]int {
        "steve": 12000,
        "jamie": 15000,
    }
	fmt.Println("personSalary map contents:", personSalary)
	for key, value := range personSalary {
        fmt.Printf("personSalary[%s] = %d\n", key, value)
	}
	delete(personSalary, "steve")
	fmt.Println("personSalary map contents:", personSalary)
	fmt.Println("length is", len(personSalary))
}