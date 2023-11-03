package main

import "fmt"

func main() {
	sliceArr := []string{"Blueberry", "Strawberries", "Watermelon", "kiwi", "mango"} //slice of containg fruit
	for _, i := range sliceArr {                                                     //retrieveing elements of slice using rang
		fmt.Print(i, " ")
	}
}
