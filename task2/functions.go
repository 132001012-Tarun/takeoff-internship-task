package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
	"unicode"

	"github.com/gorilla/mux"
)

// Implementing User interface for Employee
func (e Employee) getUsername() string { // e is a receiver
	return e.Username
}

func (e Employee) getRole() string {
	return e.Role
}

// isAdmin checks if the logged-in user is an admin
func isAdmin(user User) bool {
	return user.getRole() == "admin" //User is an interface
}
func generateNewEmployeeID() int {
	// Find the maximum ID among existing employees
	maxID := 0
	for _, emp := range employees {
		if emp.ID > maxID {
			maxID = emp.ID
		}
	}

	// Generate a new ID by incrementing the maximum ID by 1
	newID := maxID + 1
	return newID
}

// addEmployee function adds a new employee to the system
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

// viewEmployeeDetails function displays the details of a specific employee
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

// updateEmployeeDetails function updates the details of a specific employee
func updateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUT /employees/{id}")
	// Write code to update an employee and send the response
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	// Find the index of the employee in the employees slice
	index := -1
	for i, emp := range employees {
		if emp.ID == id {
			index = i
			break
		}
	}

	// Check if employee exists
	if index == -1 {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	// Parse the request body to get the updated employee data
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	// Convert the request body JSON to an Employee struct
	var updatedEmployee Employee
	err = json.Unmarshal(body, &updatedEmployee)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Validate the updated employee fields
	if !isValidName(updatedEmployee.Name) {
		http.Error(w, "Invalid name", http.StatusBadRequest)
		return
	}
	if !isValidEmail(updatedEmployee.Email) {
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}
	if !isValidPhoneNumber(updatedEmployee.Phone) {
		http.Error(w, "Invalid phone number", http.StatusBadRequest)
		return
	}
	if !isValidDateOfBirth(updatedEmployee.Birthdate) {
		http.Error(w, "Invalid date of birth", http.StatusBadRequest)
		return
	}

	// Update the employee data
	employees[index].Username = updatedEmployee.Username
	employees[index].Password = updatedEmployee.Password
	employees[index].Role = updatedEmployee.Role
	employees[index].Name = updatedEmployee.Name
	employees[index].Email = updatedEmployee.Email
	employees[index].Phone = updatedEmployee.Phone
	employees[index].Birthdate = updatedEmployee.Birthdate

	err = saveEmployeesToCSV("employees.csv", employees)
	if err != nil {
		http.Error(w, "Error saving employees", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(employees[index])
	if err != nil {
		http.Error(w, "Failed to update employee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// deleteEmployee function deletes a specific employee from the system
func deleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE /employees/{id}")
	// Get the employee ID from the URL parameter
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	// Find the index of the employee in the employees slice
	index := -1
	for i, emp := range employees {
		if emp.ID == id {
			index = i
			break
		}
	}

	// Check if employee exists
	if index == -1 {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	// Remove the employee from the employees slice
	employees = append(employees[:index], employees[index+1:]...)

	// Save the updated employees data to the CSV file
	err = saveEmployeesToCSV("employees.csv", employees)
	if err != nil {
		http.Error(w, "Error saving employees", http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Employee with ID %d deleted", id)
}

// listUpcomingBirthdays function prints the employees who have upcoming birthdays in the current month
func isValidEmail(email string) bool {
	// Use a regular expression pattern to validate the email address format
	// This pattern checks for a basic email format, but it may not catch all possible valid email addresses
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Use the MatchString function from the regexp package to check if the email matches the pattern
	match, _ := regexp.MatchString(pattern, email)

	return match
}
func isValidPhoneNumber(phoneNumber string) bool {
	// Remove any non-digit characters from the phone number
	normalizedNumber := ""
	for _, char := range phoneNumber {
		if unicode.IsDigit(char) {
			normalizedNumber += string(char)
		}
	}

	// Check the length of the normalized number
	if len(normalizedNumber) != 10 {
		return false
	}

	return true
}
func isValidName(name string) bool {
	for _, char := range name {
		if !unicode.IsLetter(char) && !unicode.IsSpace(char) {
			return false
		}
	}
	return true
}
func isValidDateOfBirth(dateOfBirth time.Time) bool {
	// Get the current date
	currentDate := time.Now().UTC()

	// Compare the date of birth with the current date
	if dateOfBirth.After(currentDate) {
		return false // Date of birth cannot be in the future
	}

	return true
}
