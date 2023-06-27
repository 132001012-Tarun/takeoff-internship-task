package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/option"
)

func addEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /employees")

	// Parse the request body to get the new employee data
	var newEmployee Employee
	err := json.NewDecoder(r.Body).Decode(&newEmployee)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	if !isValidName(newEmployee.Name) {
		http.Error(w, "Invalid name", http.StatusBadRequest)
	}
	if !isValidEmail(newEmployee.Email) {
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}

	// Validate phone number
	if !isValidPhoneNumber(newEmployee.Phone) {
		http.Error(w, "Invalid phone number", http.StatusBadRequest)
		return
	}

	// Validate date of birth
	if !isValidDateOfBirth(newEmployee.Birthdate) {
		http.Error(w, "Invalid date of birth", http.StatusBadRequest)
		return
	}
	for _, emp := range employees {
		if emp.Username == newEmployee.Username {
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newEmployee.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	newEmployee.Password = string(hashedPassword)

	// Generate a new ID for the employee
	newEmployee.ID = generateNewEmployeeID()

	// Add the new employee to the employees slice
	employees = append(employees, newEmployee)

	// Save the updated employees data to the CSV file
	err = saveEmployeesToCSV("employees.csv", employees)
	if err != nil {
		http.Error(w, "Failed to save employees", http.StatusInternalServerError)
		return
	}

	// Convert the new employee to JSON
	ctx := context.Background()
	projectID := "agile-earth-391106"
	credentialsFilePath := "/Users/tarunbyreddi/Downloads/agile-earth-391106-83b8cf42e93d.json"
	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(credentialsFilePath))
	if err != nil {
		http.Error(w, "Failed to create Firestore client", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Add the new employee document to Firestore
	docRef := client.Collection("employees").Doc(strconv.Itoa(newEmployee.ID))
	_, err = docRef.Set(ctx, newEmployee)
	if err != nil {
		http.Error(w, "Failed to save employee data to Firestore", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(newEmployee)
	if err != nil {
		http.Error(w, "Failed to create employee", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Send the response
	w.Write(jsonData)
}
