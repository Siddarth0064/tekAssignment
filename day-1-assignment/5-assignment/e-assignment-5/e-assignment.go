package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numbers []int
	i := 0
	for i <= 5 {
		randomnum := rand.Intn(100) // this line adding random numbers to slice
		numbers = append(numbers, randomnum)
		i++
	}
	fmt.Println(numbers)
	max(numbers) // this is used to pass slic to func max
}
func max(arr []int) { // this func used to find maximum numbers in the slice
	maxNum := 0
	for _, v := range arr {
		if v > maxNum {
			maxNum = v
		}
	}
	fmt.Println(maxNum)
}
