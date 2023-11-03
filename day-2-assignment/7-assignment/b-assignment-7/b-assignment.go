package main

import "fmt"

type Rectangle struct { // this is a struct of rectangle
	Width  int
	Height int
}

func (r Rectangle) calculate(h, w int) { // this is a func which accepts rectangle struct as a perimeter
	area := w * h                   //this is the formula for calculate a area of rectangle
	var perimeter int = 2 * (w + h) // this is the formula for calculate the perimeter of the racangle
	fmt.Println("the area of the rectangle is: ", area)
	fmt.Println("the perimeter of the rectangle is: ", perimeter)
}
func main() {
	rec := Rectangle{ // this is a instance of rectangle giving fields of rectangle
		Width:  24,
		Height: 54,
	}
	rec.calculate(rec.Height, rec.Width)
}
