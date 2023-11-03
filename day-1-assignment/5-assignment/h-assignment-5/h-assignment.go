package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var number []int
	i := 1
	for i <= 8 {
		randNum := rand.Intn(50) // this line adding random numbers to slice
		number = append(number, randNum)
		i++
	}
	fmt.Println(number)
	count(number) // passing slice to func to count even and odd numbers

}
func count(arr []int) {
	evenCount := 0 // initial even count value is zero
	oddCount := 0  //  initial odd count value is zero
	for _, v := range arr {

		if v%2 == 0 {
			evenCount++ // incrementing count of even numbers
		}
		if v%2 != 0 {
			oddCount++ //incrementing count of odd numbers
		}

	}
	fmt.Println("the even number count in slice is: ", evenCount)
	fmt.Println("the odd number count in slice is: ", oddCount)
}
