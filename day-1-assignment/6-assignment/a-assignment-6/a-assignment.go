package main

import "fmt"

func main() {
	studentGrades := make(map[string]float64) // creating a map as a studentGrade
	studentGrades["siddarth mp"] = 7.9        // adding a elements to map
	studentGrades["somashekar"] = 8.9         // adding a elements to map
	studentGrades["sharath"] = 6.8            // adding a elements to map
	for k, v := range studentGrades {         // retrieve the elements from map
		fmt.Println(k, " grade is ", v)
	}
	delete(studentGrades, "siddarth mp") // deleting the elements from map
	fmt.Println("after deleting elements from studentGrades is")
	for k, v := range studentGrades { // retrieve the elements from map
		fmt.Println(k, " grade is ", v)
	}
}
