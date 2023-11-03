//d. Create a group of variables using var () variable declaration syntax, assign it some values
package main

import "fmt"

func main() {
	var ( //  group of variables
		name string
		age  int
	)
	name = "siddarth"
	age = 22
	fmt.Printf("The name is %s and age is %d", name, age)

}
