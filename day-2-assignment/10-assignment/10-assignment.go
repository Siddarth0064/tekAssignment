package main

import "fmt"

func sum(nums ...int) int { // this is sum func which pass multiple parameter at a time
	total := 0
	for _, v := range nums {
		total = v + total
	}
	return total
}
func main() {
	arr1 := []int{10, 20, 30, 40, 50} // this is a slice of integer type
	fmt.Println("the slice is :", arr1)
	fmt.Println("the sum of the elements of the slice: ", sum(arr1...))
}
