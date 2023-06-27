package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

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
	ctx := context.Background()
	projectID := "agile-earth-391106"
	credentialsFilePath := "/Users/tarunbyreddi/Downloads/agile-earth-391106-83b8cf42e93d.json"
	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(credentialsFilePath))
	if err != nil {
		http.Error(w, "Failed to create Firestore client", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Delete the employee document from Firestore
	docRef := client.Collection("employees").Doc(strconv.Itoa(id))
	_, err = docRef.Delete(ctx)
	if err != nil {
		http.Error(w, "Failed to delete employee data from Firestore", http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Employee with ID %d deleted", id)
}
