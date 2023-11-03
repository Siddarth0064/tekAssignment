//b. Create a package called temperature.
//Write a program/function which accepts temperature in Fahrenheit & converts Fahrenheit to
//Celsius.

// this the main package
package main

import (
	"b-assignment-3/temperature"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("enter the temperature value ")
	var temp string
	fmt.Scan(&temp)
	v, err := strconv.Atoi(temp) // it's converts string to int
	if err != nil {              // this line check user enter int type or not and gives the err if the user enters string values
		fmt.Println("please enter the correct value")
	}
	fmt.Println("The Fahrenheit of :", temp, "°F", "  is converted to Degree Celsius is", temperature.Converts(v), "°C")

}
