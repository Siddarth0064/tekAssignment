package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 20, 10} //slice containes duplicate elements
	fmt.Println("before remove duplicate values: ", numbers)
	duplicateRemove(numbers)
}
func duplicateRemove(arr []int) { //this func used to remove duplicates elements
	uniquemap := make(map[int]bool) //using map becz keys are unique
	uniqueSLice := []int{}
	for _, v := range arr {
		if !uniquemap[v] {
			uniquemap[v] = true
			uniqueSLice = append(uniqueSLice, v)
		}
	}
	fmt.Println("after removed the duplicates values :", uniqueSLice)
}
