package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numbers []int
	i := 1
	for i <= 6 {
		randNum := rand.Intn(50) // adding random numbers to slice
		numbers = append(numbers, randNum)
		i++
	}
	fmt.Println("before double the slice : ", numbers)
	double(numbers) // passing slice to double func
	fmt.Println("after the double the slice : ", numbers)

}
func double(arr []int) { // this func used to double the elements in the slice

	for i, v := range arr {
		arr[i] = 2 * v
	}

}
