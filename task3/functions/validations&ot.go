package main

import (
	"regexp"
	"time"
	"unicode"
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

// viewEmployeeDetails function displays the details of a specific employee

// updateEmployeeDetails function updates the details of a specific employee

// deleteEmployee function deletes a specific employee from the system

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
		if !unicode.IsLetter(char) && !unicode.IsSpace(char) { // check wheather if a letter is not char
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
