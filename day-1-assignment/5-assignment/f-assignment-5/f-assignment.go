package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50} //this is a slice which contains 5 diffrentrs elements
	var num int                          // this is user wants to search in the slice
	fmt.Println("enter the number to search from slice: ")
	fmt.Scan(&num)
	searchNum(numbers, num) // this line passing slice and number to find in the slice
}
func searchNum(arr []int, x int) { // this is a func to find specifc elements in the slice
	for _, v := range arr {
		if v == x {
			fmt.Println("number is found in the slice")
			return
		}
	}
	fmt.Println("number is not found in the slice")
}
