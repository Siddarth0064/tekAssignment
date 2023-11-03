package main

import (
	"c-assignment-8/model"
	"c-assignment-8/person"
)

func main() {
	p := model.Person{ // this is the instance of person struct and giving values to i
		Name: "siddarth",
		Age:  22,
	}

	person.PrintPerson(p)
}
