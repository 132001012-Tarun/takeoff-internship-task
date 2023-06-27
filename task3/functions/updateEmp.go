package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

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

	// Retrieve the current employee details from Firestore
	ctx := context.Background()
	projectID := "agile-earth-391106"
	credentialsFilePath := "/path/to/your/credentials.json"

	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(credentialsFilePath))
	if err != nil {
		http.Error(w, "Failed to create Firestore client", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	docRef := client.Collection("employees").Doc(strconv.Itoa(id))
	docSnapshot, err := docRef.Get(ctx)
	if err != nil {
		http.Error(w, "Failed to retrieve employee details from Firestore", http.StatusInternalServerError)
		return
	}

	// Parse the current employee details from Firestore
	var currentEmployee Employee
	err = docSnapshot.DataTo(&currentEmployee)
	if err != nil {
		http.Error(w, "Failed to parse employee details from Firestore", http.StatusInternalServerError)
		return
	}

	// Parse the form values to get the updated employee data
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form values", http.StatusBadRequest)
		return
	}

	// Update the employee data
	currentEmployee.Username = r.FormValue("username")
	currentEmployee.Password = r.FormValue("password")
	currentEmployee.Role = r.FormValue("role")
	currentEmployee.Name = r.FormValue("name")
	currentEmployee.Email = r.FormValue("email")
	currentEmployee.Phone = r.FormValue("phone")

	// Parse and validate the updated birthdate
	birthdateStr := r.FormValue("birthdate")
	birthdate, err := time.Parse("2006-01-02T15:04:05Z", birthdateStr)
	if err != nil {
		http.Error(w, "Invalid birthdate format", http.StatusBadRequest)
		return
	}
	currentEmployee.Birthdate = birthdate

	// Save the updated employee data to Firestore
	_, err = docRef.Set(ctx, currentEmployee)
	if err != nil {
		http.Error(w, "Failed to update employee details in Firestore", http.StatusInternalServerError)
		return
	}

	// Save the updated employee data to the CSV file

	for i, emp := range employees {
		if emp.ID == id {
			employees[i] = currentEmployee
			break
		}
	}

	err = saveEmployeesToCSV("employees.csv", employees)
	if err != nil {
		http.Error(w, "Failed to save employees to CSV", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(currentEmployee)
	if err != nil {
		http.Error(w, "Failed to marshal updated employee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
