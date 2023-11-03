package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Grade string
}

var DataBase = make(map[int]Student)

func main() {
	Student1 := Student{Id: 10001, Name: "Sangamesh", Age: 22, Grade: "A+"}
	Student2 := Student{Id: 10002, Name: "Somashekar", Age: 23, Grade: "A"}
	DataBase[10001] = Student1
	DataBase[10002] = Student2

	http.HandleFunc("/students", StudentsHandler)
	http.ListenAndServe(":8080", nil)
}

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check the HTTP method
	if r.Method == http.MethodGet {
		// Check if it's a request for a specific student
		studentIDStr := r.URL.Query().Get("studentId")
		if studentIDStr != "" {
			studentID, err := strconv.Atoi(studentIDStr)
			if err != nil {
				errResponse := map[string]string{"error": "Invalid student ID"}
				jsonErr, err := json.Marshal(errResponse)
				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintln(w, "Internal Server Error")
					return
				}
				w.WriteHeader(http.StatusBadRequest)
				w.Write(jsonErr)
				return
			}

			student, found := DataBase[studentID]
			if !found {
				errResponse := map[string]string{"error": "Student not found"}
				jsonErr, err := json.Marshal(errResponse)
				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintln(w, "Internal Server Error")
					return
				}
				w.WriteHeader(http.StatusNotFound)
				w.Write(jsonErr)
				return
			}

			// Return the student as JSON
			studentJSON, err := json.Marshal(student)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "Internal Server Error")
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(studentJSON)
		} else {
			// Return the entire student map as JSON
			studentMapJSON, err := json.Marshal(DataBase)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "Internal Server Error")
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(studentMapJSON)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not allowed")
	}
}
