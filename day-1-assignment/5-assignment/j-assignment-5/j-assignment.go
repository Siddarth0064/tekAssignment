package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var slice1 []int
	var slice2 []int
	var slice3 []int
	i := 1
	for i <= 3 {
		randNum := rand.Intn(20) // adding random numbers to slice 1
		slice1 = append(slice1, randNum)
		i++
	}
	i = 1
	for i <= 3 { // adding random numbers to slice 2
		randNum := rand.Intn(50)
		slice2 = append(slice2, randNum)
		i++
	}
	fmt.Println("before concatenation....")
	fmt.Println("slice 1 : ", slice1)
	fmt.Println("slice 2 :", slice2)
	fmt.Println("after concatenation....")

	slice3 = append(slice3, slice1...) //appending
	slice3 = append(slice3, slice2...)
	fmt.Println("adding two diffrent slice to new slice.   :  ", slice3)
	slice1 = append(slice1, slice2...)
	fmt.Println("adding slice 2 to slice 1. now slice 1 is : ", slice1)

}
