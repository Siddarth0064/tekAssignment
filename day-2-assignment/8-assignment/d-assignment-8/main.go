package main

import (
	"d-assignment-8/model"
	"d-assignment-8/shape"
	"fmt"
)

func main() {
	r := model.Circle{ //creating a instance variable of circle
		Radius: 24,
	}
	fmt.Println(shape.Area(r))
}
