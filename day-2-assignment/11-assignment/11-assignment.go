package main

import "fmt"

type Address struct { //this is a address struct which contains street,city and zipcode
	Street  string
	City    string
	ZipCode int
}
type Person struct { // this is a person struct which contains name and address struct
	Name string
	Address
}

func main() {
	p := Person{ // this is a instance of person
		Name: "siddarth",
		Address: Address{Street: "vijaynagara", // this is a instance of address
			City:    "bangalore",
			ZipCode: 571101},
	}
	fmt.Printf("the name is %v and address of street is %v and city is %v and zipcode is %v", p.Name, p.Street, p.City, p.ZipCode)
}
