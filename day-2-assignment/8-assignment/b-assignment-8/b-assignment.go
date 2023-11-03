package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num string //crating string varible
	fmt.Println("enter the number to get factorials")
	fmt.Scan(&num)
	v, err := strconv.Atoi(num) //converting string to int
	if err != nil {             //checking err is occured or not
		fmt.Println("please enter the number")
		return
	}
	fact := factories(v) // passing value to factories func and storingn to varibales
	fmt.Println(,"the factorials of number is :", fact)
}

func factories(x int) int { //this is recurssion factories func to calculate factories
	if x == 0 {
		return 1
	}
	return x * factories(x-1)
}
