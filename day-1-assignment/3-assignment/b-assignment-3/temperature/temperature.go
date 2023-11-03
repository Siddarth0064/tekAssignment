//b. Create a package called temperature.
// Write a program/function which accepts temperature in Fahrenheit & converts Fahrenheit to
// Celsius.

// this the temperature package
package temperature

//import "fmt"

func Converts(x int) float64 { // this function converts the temperature Fahrenheit to Degree Celsius
	celsius := float64(x-32) * 5 / 9
	return celsius
}
