package main

import ( // libraries

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// isValidName checks if a name contains only characters
func main() {
	// Add some dummy employees for testing     .///basically our codes start excecuting from here

	loadedEmployees, err := loadEmployeesFromCSV("employees.csv") // loads data if it exists in employees.csv file otherwise returns error
	if err != nil {
		fmt.Println("Error loading employees:", err)
		return
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
