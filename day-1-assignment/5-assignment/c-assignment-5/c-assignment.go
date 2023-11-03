package main

import "fmt"

func main() {
	var numbers []int // creating empty slice
	i := 5
	for i <= 25 { // adding elements to empty slice using loop by increamenting 5
		numbers = append(numbers, i)
		i = i + 5
		fmt.Println(numbers)
	}
	fmt.Println("the length of the slice is :", len(numbers))   // this line gives length of the slice
	fmt.Println("the capacity of the slice is: ", cap(numbers)) //this line gives capavity of the slice
	numbers = append(numbers[:2], numbers[3:]...)               // this line removing specific elements from slice
	fmt.Println(numbers)
}
