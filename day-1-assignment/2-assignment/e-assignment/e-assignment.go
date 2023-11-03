// e. Create a group of constants using var () group constant declaration syntax, assign it some
// values & print it. Try to change its values & check if there is some error.
package main

import (
	"fmt"
)

func main() {
	const ( // group of constant variables
		name = "siddarth"
		age  = 22
	)
	fmt.Printf("The name is %s and age is %d", name, age)

}
