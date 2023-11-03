package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50} //array of contains 5 elements
	for i := 0; i < len(arr); i++ {   // using for loop to retrieve elements
		fmt.Print(arr[i], " ")
	}
}
