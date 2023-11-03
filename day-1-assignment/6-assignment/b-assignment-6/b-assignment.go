package main

import "fmt"

func main() {
	fruits := make(map[string][]string) //creating map as a fruit which stroes key as a int and value as a string of slices
	fruits["mango"] = []string{" A round fruit with a red color", "its a desserts.", "costliest"}
	fruits["banana"] = []string{"a longest fruit", "white color"}
	fruits["apple"] = []string{" fresh and well formed ", "its helps to hydrate skin"}
	fruits["orange"] = []string{"a rounded fruit", "orange color", "its helps to hydrate"}
	for k, v := range fruits {
		for _, v1 := range v { // retriev the elements from the map
			fmt.Print(k, ":  ", v1)
		}
		fmt.Println()
	}
}
