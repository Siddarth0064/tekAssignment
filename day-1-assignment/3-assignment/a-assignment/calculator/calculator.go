//a. Create a package named calculator.
//Inside the calculator package, define a below functions that

// this is a package of calculator
package calculator

func Square(x int) int { // this function calculate Square of the number
	return x * x // and return it
}
func Sum(x, y int) int { // this function calculate Sum of two number
	return x + y //and return it
}
func Difference(x, y int) int { // this function calculate Difference of two number
	return x - y // and return it
}
func Product(x, y int) int { // this function calculate Product of two number
	return x * y // and return it
}
func QuotRemaind(x, y int) (int, int) { // this function calculate Remainder and Quotient of two number
	return (x / y), (x % y) // and return Remainder and Quotient value
}
