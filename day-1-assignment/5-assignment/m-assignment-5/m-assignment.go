package main

import (
	"fmt"
)

func main() {
	slice1 := []int{10, 20, 30, 40, 50}
	slice2 := []int{10, 20, 30, 40, 5}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] { // this is for equating the elements from two diffrent slice
			fmt.Println("slice are not equals")
			return
		}
	}
	fmt.Println("slice are equals")
}
