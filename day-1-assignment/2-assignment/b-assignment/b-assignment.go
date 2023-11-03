//b. Declare a variable of type float64 and assign it a value of 3.14. Print the variable.
package main

import "fmt"

var f_num1 float64 = 3.14 // its a global variable

func main() {
	var f_num2 float64 = 3.14 // its a local variable using var keyword
	f_num3 := 3.14            // its a local variable using short declaration operatoscope
	fmt.Println("the variable from global scope f-num1: ", f_num1)
	fmt.Println("the variable from local scope f-num2: ", f_num2)
	fmt.Println("the variable from local scope f-num3: ", f_num3)

}
