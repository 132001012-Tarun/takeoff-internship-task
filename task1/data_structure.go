package main

import "time"

type User interface {
	getUsername() string
	getRole() string
}

// Employee struct represents an employee record
type Employee struct {
	ID        int
	Username  string
	Password  string
	Role      string
	Name      string
	Email     string
	Phone     string
	Birthdate time.Time
}

var employees []Employee // declaring employees as slice of Employee struct
