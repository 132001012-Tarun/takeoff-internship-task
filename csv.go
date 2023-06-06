package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

// saveEmployeesToCSV saves the employee records to a CSV file
func saveEmployeesToCSV(filename string, employees []Employee) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"ID", "Username", "Password", "Role", "Name", "Email", "Phone", "Birthdate"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write employee records
	for _, employee := range employees {
		record := []string{
			strconv.Itoa(employee.ID),
			employee.Username,
			employee.Password,
			employee.Role,
			employee.Name,
			employee.Email,
			employee.Phone,
			employee.Birthdate.Format("2006-01-02"),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

// loadEmployeesFromCSV loads the employee records from a CSV file
func loadEmployeesFromCSV(filename string) ([]Employee, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	employees := make([]Employee, 0, len(records)-1)
	for i := 1; i < len(records); i++ {
		record := records[i]
		id, _ := strconv.Atoi(record[0])
		birthdate, _ := time.Parse("2006-01-02", record[7])

		employee := Employee{
			ID:        id,
			Username:  record[1],
			Password:  record[2],
			Role:      record[3],
			Name:      record[4],
			Email:     record[5],
			Phone:     record[6],
			Birthdate: birthdate,
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

// Load employees from CSV file
//loadedEmployees, err := loadEmployeesFromCSV("employees.csv")
//if err != nil {
//	fmt.Println("Error loading employees:", err)
//	return
//}

// Print loaded employees
//for _, employee := range loadedEmployees {
//	fmt.Println(employee)
//}
