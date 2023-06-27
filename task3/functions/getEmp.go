package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getAllEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /employees")
	// Retrieve all employees
	employeesData := make([]Employee, len(employees))
	copy(employeesData, employees)

	// Convert employeesData to JSON
	jsonData, err := json.Marshal(employeesData)
	if err != nil {
		http.Error(w, "Failed to retrieve employees", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Send the response
	w.Write(jsonData)
}
