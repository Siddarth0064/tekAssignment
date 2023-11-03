// a. Implement a method called checkEvenOdd , it takes int as an input.
// check if it is even or odd.
package main

import (
	"fmt"
	"strconv"
)

func checkEvenOdd(x int) { // this function is check its even or odd
	if x%2 == 0 { //this condition is check only even number and return from this function
		fmt.Println("number is Even :", x)
		return // if its even it will return no more execution
	}

	fmt.Println("number is Odd :", x) // this line gives only Odd numbers

}
func main() {
	fmt.Println("enter the number to check its even or odd")
	var num string
	fmt.Scan(&num)
	v, err := strconv.Atoi(num) // this line check user enter only interge or not if user enyers string datatype it will return
	if err != nil {             // this conditon throw an error if user enter the string data type
		fmt.Println("plese recheck and enter the number")
		return // if error is occured it will return
	}
	checkEvenOdd(v) // this line passing the parameter to checkEvenOdd func

}
