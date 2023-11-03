package main

import (
	"fmt"
	"math"
)

type Shape interface { //this is a shape interface which contains area method
	Area() float32
}
type Circle struct { // this is a circle struct which contains radius
	Radius int
}
type Rectangle struct { // this is a rectangle struct which contains length and height
	Length int
	Height int
}

func (c Circle) Area() float64 { //this is a func of area which accepts circle struct
	ar := math.Pi * float64(c.Radius) * float64(c.Radius)
	return ar
}
func (r Rectangle) Area() float64 { //this is a func of area which accepts rectangle struct
	arr := r.Length * r.Height
	return float64(arr)
}
func main() {
	cr := Circle{ //this is a instance of circle
		Radius: 56,
	}
	rec := Rectangle{ // this is a instance of rectangle
		Length: 47,
		Height: 49,
	}
	fmt.Println("the area of the circle is : ", Circle.Area(cr))
	fmt.Println("the area of the rectangel is : ", Rectangle.Area(rec))
}
