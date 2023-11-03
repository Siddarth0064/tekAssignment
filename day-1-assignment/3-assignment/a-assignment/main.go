//a. Create a package named calculator.
//Inside the calculator package, define a below functions that

package main //this is main package

import (
	"a-assignment/calculator" // retrieving the calculator package
	"fmt"
)

func main() {
	a := 20
	b := 4
	fmt.Println(calculator.Difference(a, b), "is the differenc of 20 and 4")           //this line gives Difference of two numbers
	fmt.Println(calculator.Product(a, b), "is the product of 20 and 4")                // this line gives Product of two numbers
	fmt.Println(calculator.Square(a), " is the square of 20")                          // this line gives Square of that number
	fmt.Println(calculator.Sum(a, b), " is the sum of 20 and 4")                       // this line gives Sum of two number
	remainder, quotient := calculator.QuotRemaind(a, b)                                // this line stores Remainder and Quotient of two numbers
	fmt.Println(remainder, quotient, " is ths quotient and remender of 20 and 4 and ") // this line prints Remainder and Quotient of two numbers
}
