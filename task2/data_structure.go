package main

import "time"

type User interface {
	getUsername() string
	getRole() string
}

// Employee struct represents an employee record
type Employee struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Birthdate time.Time `json:"birthdate"`
}

var employees []Employee // declaring employees as slice of Employee struct
