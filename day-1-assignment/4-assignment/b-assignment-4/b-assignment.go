package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	fmt.Println("welcome!.. to the NUMBER GUESS GAME..")
	randomnumber := rand.Intn(10) //its generate random numbers
	var num string
	for {
		fmt.Println("enter the your number with in 1 to 10")
		fmt.Scan(&num)
		v, err := strconv.Atoi(num) //it will check user enter integer data type or not
		if err != nil {
			fmt.Println("please enter only numbers")
			continue
		}

		if randomnumber == v { // its gives guess is correct or not
			fmt.Println("Congratulations... WOwwww.  your guess is correct:", v)
			break
		}
		if randomnumber > v { //its gives user guess number is too low worning signals
			fmt.Println("sorry!.. your guess is too low")
			continue
		}
		if randomnumber < v { ////its gives user guess number is too high  worning signals
			fmt.Println("sorry!.. your guess is too high")
			continue
		}
	}
}
