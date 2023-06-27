package main

import ( // libraries

	"context"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/api/option"
)

// isValidName checks if a name contains only characters
func main() {
	// Add some dummy employees for testing     .///basically our codes start excecuting from here
	ctx := context.Background()
	projectID := "agile-earth-391106"
	credentialsFilePath := "/Users/tarunbyreddi/Downloads/agile-earth-391106-83b8cf42e93d.json"
	csvFilePath := "employees.csv"

	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(credentialsFilePath))
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer client.Close()

	loadedEmployees, err := loadEmployeesFromCSV("employees.csv") // loads data if it exists in employees.csv file otherwise returns error
	if err != nil {
		fmt.Println("Error loading employees:", err)
		return
	}
	err = migrateCSVToFirestore(ctx, client, csvFilePath)
	if err != nil {
		log.Fatalf("Failed to migrate CSV to Firestore: %v", err)
	}

	fmt.Println("Employee Management System")
	fmt.Println("--------------------------")
	employees = loadedEmployees // stores the loaded data into a csv file
	fmt.Println()
	router := mux.NewRouter()                                              // server initialization
	router.HandleFunc("/employees", getAllEmployeesHandler).Methods("GET") // handling request
	router.HandleFunc("/employees/{id}", deleteEmployeeHandler).Methods("DELETE")
	router.HandleFunc("/employees", addEmployeeHandler).Methods("POST")
	router.HandleFunc("/employees/{id}", updateEmployeeHandler).Methods("PUT")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router)) // server started

}

// MigrateCSVToFirestore migrates data from a CSV file to Firestore
func migrateCSVToFirestore(ctx context.Context, client *firestore.Client, csvFilePath string) error {
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %v", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	employeeData, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %v", err)
	}

	// Iterate through each row in the CSV file
	for _, row := range employeeData {
		id, _ := strconv.Atoi(row[0])
		username := row[1]
		password := row[2]
		role := row[3]
		name := row[4]
		email := row[5]
		phone := row[6]
		birthdate, _ := time.Parse("2006-01-02", row[7])
		// ... parse other fields

		employee := Employee{
			ID:        id,
			Username:  username,
			Password:  password,
			Role:      role,
			Name:      name,
			Email:     email,
			Phone:     phone,
			Birthdate: birthdate,
			// ... set other fields
		}

		// Save the employee data to Firestore
		docRef := client.Collection("employees").Doc(strconv.Itoa(employee.ID))
		_, err = docRef.Set(ctx, employee)
		if err != nil {
			return fmt.Errorf("failed to save employee data to Firestore: %v", err)
		}
	}

	return nil
}

