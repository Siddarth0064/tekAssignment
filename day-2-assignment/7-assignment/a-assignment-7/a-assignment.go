package main

import "fmt"

type Employee struct { //this is a employee struct which contains fields of employees
	EmpId   int
	EmpName string
	EmpAge  int
	EmpCity string
}

func main() {
	emp1 := Employee{ // this is a instance of employe of  1
		EmpId:   1001,
		EmpName: "siddarth mp",
		EmpAge:  22,
		EmpCity: "mysore",
	}
	emp2 := Employee{ // this is a instance of employee of  2
		EmpId:   1002,
		EmpName: "somashekar",
		EmpAge:  23,
		EmpCity: "bangalore",
	}
	fmt.Printf("EmpId: %v , EmpName: %v ,EmpAge: %v , EmpCity: %v \n", emp1.EmpId, emp1.EmpName, emp1.EmpAge, emp1.EmpCity)
	fmt.Printf("EmpId: %v , EmpName: %v ,EmpAge: %v , EmpCity: %v \n", emp2.EmpId, emp2.EmpName, emp2.EmpAge, emp2.EmpCity)

}
