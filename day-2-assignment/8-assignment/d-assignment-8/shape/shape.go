package shape

import (
	"d-assignment-8/model"
	"math"
)

func Area(c model.Circle) float64 {
	return math.Pi * c.Radius * c.Radius
}
