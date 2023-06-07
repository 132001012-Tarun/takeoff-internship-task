package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

func isValidName(name string) bool {
	for _, char := range name {
		if !unicode.IsLetter(char) && !unicode.IsSpace(char) {
			return false
		}
	}
	return true
}

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

// addEmployee function adds a new employee to the system
func addEmployee() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Add Employee")

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n') // take the input untill we press enter
	username = strings.TrimSpace(username) //TrimSpace will remove leading and trailing whitespace characters from a string

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	fmt.Print("Role: ")
	role, _ := reader.ReadString('\n')
	role = strings.TrimSpace(role)

	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if !isValidName(name) {
		fmt.Println("Invalid name! Name should only contain alphabetic characters.")
		return
	}

	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Phone: ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	fmt.Print("Birthdate (yyyy-mm-dd): ")
	birthdateStr, _ := reader.ReadString('\n')
	birthdateStr = strings.TrimSpace(birthdateStr)
	birthdate, _ := time.Parse("2006-01-02", birthdateStr)

	employee := Employee{
		ID:        len(employees) + 1,
		Username:  username,
		Password:  password,
		Role:      role,
		Name:      name,
		Email:     email,
		Phone:     phone,
		Birthdate: birthdate,
	}

	employees = append(employees, employee)
	err := saveEmployeesToCSV("employees.csv", employees)
	if err != nil {
		fmt.Println("Error saving employees:", err)
		return
	}
	fmt.Println("Employee added successfully!")
}

// viewEmployeeDetails function displays the details of a specific employee
func viewEmployeeDetails(user User) {
	username := user.getUsername()
	for _, employee := range employees {
		if employee.Username == username {
			fmt.Println("Employee Details:")
			printEmployeeDetails(employee)
			return
		}
	}
	fmt.Println("Employee not found!")
}

// updateEmployeeDetails function updates the details of a specific employee
func updateEmployeeDetails(Currentuser User) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Update Employee Details")
	if isAdmin(Currentuser) { // if you are in admin side you can modify anyone's details
		fmt.Print("Enter the username of the employee to update: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)
		for i, employee := range employees {
			if employee.Username == username {
				fmt.Println("Enter new details:")

				fmt.Print("Name: ")
				name, _ := reader.ReadString('\n')
				name = strings.TrimSpace(name)

				if !isValidName(name) {
					fmt.Println("Invalid name! Name should only contain alphabetic characters.")
					return
				}

				fmt.Print("Email: ")
				email, _ := reader.ReadString('\n')
				email = strings.TrimSpace(email)

				fmt.Print("Phone: ")
				phone, _ := reader.ReadString('\n')
				phone = strings.TrimSpace(phone)

				employees[i].Name = name
				employees[i].Email = email
				employees[i].Phone = phone

				err := saveEmployeesToCSV("employees.csv", employees)
				if err != nil {
					fmt.Println("Error saving employees:", err)
					return
				}

				fmt.Println("Employee details updated successfully!")
				return
			}
		}
		fmt.Println("Employee not found!")
	} else { // if you are not admin
		username := Currentuser.getUsername()
		for i, employee := range employees {
			if employee.Username == username {
				fmt.Println("Enter new details:")

				fmt.Print("Name: ")
				name, _ := reader.ReadString('\n')
				name = strings.TrimSpace(name)

				if !isValidName(name) {
					fmt.Println("Invalid name! Name should only contain alphabetic characters.")
					return
				}

				fmt.Print("Email: ")
				email, _ := reader.ReadString('\n')
				email = strings.TrimSpace(email)

				fmt.Print("Phone: ")
				phone, _ := reader.ReadString('\n')
				phone = strings.TrimSpace(phone)

				employees[i].Name = name
				employees[i].Email = email
				employees[i].Phone = phone

				err := saveEmployeesToCSV("employees.csv", employees)
				if err != nil {
					fmt.Println("Error saving employees:", err)
					return
				}

				fmt.Println("Employee details updated successfully!")
				return
			}
		}
		fmt.Println("Employee not found!")
	}

}

// deleteEmployee function deletes a specific employee from the system
func deleteEmployee() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Delete Employee")

	fmt.Print("Enter username of the employee to delete: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	for i, employee := range employees {
		if employee.Username == username {
			employees = append(employees[:i], employees[i+1:]...)
			err := saveEmployeesToCSV("employees.csv", employees)
			if err != nil {
				fmt.Println("Error saving employees:", err)
				return
			}
			fmt.Println("Employee deleted successfully!")
			return
		}
	}
	fmt.Println("Employee not found!")
}

// listEmployees function prints the details of all employees
func listEmployees() {
	fmt.Println("Employee List:")
	for _, employee := range employees {
		printEmployeeDetails(employee)
	}
}

// listUpcomingBirthdays function prints the employees who have upcoming birthdays in the current month
func listUpcomingBirthdays() {
	currentMonth := time.Now().Month()
	fmt.Println("Employees with Birthdays on this Month:")

	for _, employee := range employees {
		if employee.Birthdate.Month() == currentMonth {
			fmt.Println("Name:", employee.Name)
			fmt.Println("Birthdate:", employee.Birthdate.Format("2006-01-02"))
			fmt.Println()
		}
	}
}

// searchEmployee function searches for an employee by name and displays all their fields except password
func searchEmployeeByName() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Search Employee")

	fmt.Print("1.Enter name to search: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if !isValidName(name) {
		fmt.Println("Invalid name! Name should only contain alphabetic characters.")
		return
	}

	found := false
	for _, employee := range employees {
		if strings.ToLower(employee.Name) == strings.ToLower(name) {
			found = true
			fmt.Println("Employee Details:")
			printEmployeeDetails(employee)
		}
	}

	if !found {
		fmt.Println("No matching employees found!")
	}
}
func searchEmployeeById() { // serach by id
	var num int
	fmt.Print("Enter the Id Of the employee: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Error reading integer:", err)
		return
	}
	found := false
	for _, employee := range employees {
		if employee.ID == num {
			found = true
			fmt.Println("Employee Details:")
			printEmployeeDetails(employee)
		}
	}

	if !found {
		fmt.Println("No matching employees found!")
	}
}
