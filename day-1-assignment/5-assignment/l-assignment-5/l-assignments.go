package main

import "fmt"

func main() {
	numbers := []int{22, 34, 45, 67, 87, 88} // the slice which contains even numbers and odd numbers
	fmt.Println("the slice is : ", numbers)
	sumofEven(numbers)
}
func sumofEven(arr []int) { // this is the func which sum the all the even numbers from the slice
	sumOfEven := 0
	for _, v := range arr {
		if v%2 == 0 {
			sumOfEven = v + sumOfEven
		}
	}
	fmt.Println("the sum of even numbers from the slice is: ", sumOfEven)
}
