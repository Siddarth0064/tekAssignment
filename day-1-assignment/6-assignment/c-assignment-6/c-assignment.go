package main

import "fmt"

func main() {
	studentData := make(map[string]map[string]interface{}) // this is a map which contains key as string and value as a another amp
	studentData["siddarth mp"] = map[string]interface{}{"Age": 22, "Grade": "A", "City": "mysore"}
	studentData["somasheker"] = map[string]interface{}{"Age": 21, "Grade": "A+", "City": "bangalore"}
	studentData["janardhan"] = map[string]interface{}{"Age": 23, "Grade": "B+", "City": "delhi"}
	for k, v := range studentData {
		fmt.Print(k, ": ")
		for k1, v1 := range v {
			fmt.Print(k1, ": ", v1, " ")
		}
		fmt.Println()
	}

}
