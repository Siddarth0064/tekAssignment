// a. Declare a variable of type int and assign it a value of 42. Print the variable.

package main

import "fmt"

var num1 int = 42 // its a global variable

func main() {
	num2 := 42        // its a local variable using short declaration operator
	var num3 int = 42 // its a local variable using var keyword
	fmt.Println("the variable from globale scope num1: ", num1)
	fmt.Println("the variable from local  scope num2: ", num2)
	fmt.Println("the varible from local scope and using var keyword num3: ", num3)

}
