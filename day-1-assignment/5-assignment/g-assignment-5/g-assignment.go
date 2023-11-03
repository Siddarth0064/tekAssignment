package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numbers []int
	i := 1
	for i <= 5 {
		randomNum := rand.Intn(100) // this line adding random numbers to slice
		numbers = append(numbers, randomNum)
		i++
	}
	fmt.Println("before revers a slice: ", numbers)
	reversSlice(numbers)
}
func reversSlice(arr []int) { // this func is used to revers the original slice
	len := len(arr)
	fmt.Print("after the revers of slice is: ")
	for i, j := 0, len-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println(arr)
}
