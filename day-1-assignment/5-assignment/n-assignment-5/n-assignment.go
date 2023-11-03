package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	var findNumber int
	fmt.Println("enter the number to find index from the slice")
	fmt.Scan(&findNumber) //this line for user enter the number to find the index
	for i, v := range numbers {
		if v == findNumber {
			fmt.Println("the number is found in the index of :", i)
			return
		}
	}
	fmt.Println("the number is not found in the slice  ")
}
