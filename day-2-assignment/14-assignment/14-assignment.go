package main

import "fmt"

func divide(numerator, denominator int) { // this is a func whic throw the panic manually
	if denominator == 0 { // if denominator is zero panic will be rise
		panic("RAISEING PANIC MANUALLY : denominator is zero")
	}
	fmt.Println("no problem in denominator")
}
func safeDivied(numerator, denominator int) { // this is a func to recovery the panic
	defer recovery()
	divide(numerator, denominator)

}
func main() {
	fmt.Println("main starting")
	safeDivied(10, 5)
	safeDivied(10, 0)
	fmt.Println("main closing normally")
}

func recovery() { // this is the recovery func which recovery from the panic
	m := recover()

	if m != nil {
		fmt.Println(m)

	}
}
