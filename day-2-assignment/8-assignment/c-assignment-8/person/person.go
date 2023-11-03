package person

import "c-assignment-8/model"
import "fmt"

func PrintPerson(p model.Person) {
	fmt.Println("the name is :", p.Name)
	fmt.Println("age is :", p.Age)
}
