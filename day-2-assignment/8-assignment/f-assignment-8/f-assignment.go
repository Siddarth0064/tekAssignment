package main

import "fmt"

type Student struct { // this is a student struct which contains name and age
	Name string
	Age  int
}

func (s *Student) update(ss int) { // this is a update func for age of student struct
	s.Age = ss + 1
}
func main() {
	s := Student{ //this is a instance of student and filling fields
		Name: "siddarth",
		Age:  23,
	}
	fmt.Println("before the upadte age is :")
	fmt.Println(s.Age)
	fmt.Println("after the upadte age is : ")
	s.update(s.Age)
	fmt.Println(s.Age)

}
