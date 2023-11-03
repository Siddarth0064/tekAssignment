// c. Create a constant called pi and assign it the value of Pi (math.Pi). Print the constant.
package main

import (
	"fmt"
	"math"
)

const Pi1 float64 = math.Pi // constant global variable

func main() {
	const pi2 = math.Pi // constant local variable

	fmt.Println("the variable from global scope pi1: ", Pi1)
	fmt.Println("the variable from local scope pi2: ", pi2)
}
