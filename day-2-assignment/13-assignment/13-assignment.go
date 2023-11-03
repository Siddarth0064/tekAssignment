package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var filename string = "sample.txt"  // creating variable and assigning file name
	v, err := ioutil.ReadFile(filename) // passing file name to func as a readfile
	if err != nil {                     // checking err is occured or not
		fmt.Println(err)
		return
	}
	str := string(v)           //this is a func of area which accepts circle struct
	arr := strings.Fields(str) // string converted to arrays and removed whitespace which meanes we can calculate the word of the string
	fmt.Println("the length of word in the smpale.txt file is: ", len(arr))
}
