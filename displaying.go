package main

import "fmt"

// login function prompts the user to enter username and password
func printEmployeeDetails(employee Employee) { // basically a normal print function
	fmt.Println("ID:", employee.ID)
	fmt.Println("Username:", employee.Username)
	fmt.Println("Role:", employee.Role)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Email:", employee.Email)
	fmt.Println("Phone:", employee.Phone)
	fmt.Println("Birthdate:", employee.Birthdate.Format("2006-01-02"))
	fmt.Println()
}
