package main

import "time"

type User interface {
	getUsername() string
	getRole() string
}

// Employee struct represents an employee record
type Employee struct {
	ID        int       `firestore:"id"`
	Username  string    `firestore:"username"`
	Password  string    `firestore:"password"`
	Role      string    `firestore:"role"`
	Name      string    `firestore:"name"`
	Email     string    `firestore:"email"`
	Phone     string    `firestore:"phone"`
	Birthdate time.Time `firestore:"birthdate"`
}

var employees []Employee // declaring employees as slice of Employee struct
