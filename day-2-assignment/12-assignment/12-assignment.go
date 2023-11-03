package main

import "fmt"

type Animal interface { // this is a animal interface which contains makesound func
	MakeSound() string
}
type Cat struct { // this is cat struct
	CatSound string
}
type Dog struct { // this is dog struct
	DogSound string
}

func (d Dog) MakeSound() string { // this is a func with refrence with dog struct
	sound := " bow bow"
	return sound
}
func (c Cat) MakeSound() string { // this is a func with refrence with cat struct
	sound := "meoww"
	return sound
}
func main() {
	c := Cat{} // creating instance of cat
	d := Dog{} // creating instance of dog
	fmt.Println("cat sound like ", c.MakeSound())
	fmt.Println("dog sound like ", d.MakeSound())

}
