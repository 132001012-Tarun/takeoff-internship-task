package main

import ( // libraries

	"fmt"
)

// isValidName checks if a name contains only characters
func main() {
	// Add some dummy employees for testing     .///basically our codes start excecuting from here

	loadedEmployees, err := loadEmployeesFromCSV("employees.csv")
	if err != nil {
		fmt.Println("Error loading employees:", err)
		return
	}
	// Login screen
	fmt.Println("Employee Management System")
	fmt.Println("--------------------------")
	employees = loadedEmployees
	user, success := login() // login function called and it returns (user,success) values .  (where user is an interface and success is a bool)

	if success { // if success is true then this will excecute
		fmt.Println("Login successful!")
		fmt.Println()

		if isAdmin(user) { // if the user returned from the login() function is admin then the below code will excecute
			// Admin operations
			for {
				fmt.Println("Admin Menu:")
				fmt.Println("1. Add Employee")
				fmt.Println("2. View Employee Details")
				fmt.Println("3. Update Employee Details")
				fmt.Println("4. Delete Employee")
				fmt.Println("5. List All Employees")
				fmt.Println("6. List Employees with Birthdays this month")
				fmt.Println("7. Search Employee by name")
				fmt.Println("8. Search Employee by Id")
				fmt.Println("9. Logout")

				var choice int
				fmt.Print("Enter your choice (1-9): ")
				fmt.Scanln(&choice)
				fmt.Println()

				switch choice {
				// these are the various operations that admin can form and based on the choice of the admin he/she can perform different tasks.
				case 1:
					addEmployee()
				case 2:
					viewEmployeeDetails(user)
				case 3:
					updateEmployeeDetails(user)
				case 4:
					deleteEmployee()
				case 5:
					listEmployees()
				case 6:
					listUpcomingBirthdays()
				case 7:
					searchEmployeeByName()
				case 8:
					searchEmployeeById()
				case 9:
					fmt.Println("Logout successful!")
					err := saveEmployeesToCSV("employees.csv", employees)
					if err != nil {
						fmt.Println("Error saving employees:", err)
						return
					}
					return
				default:
					fmt.Println("Invalid choice! Please try again.")
				}

				fmt.Println()
			}
		} else {
			// Non-admin operations
			for {
				fmt.Println("User Menu:")
				fmt.Println("1. View My Details")
				fmt.Println("2. Update My Details")
				fmt.Println("3. Logout")

				var choice int
				fmt.Print("Enter your choice (1-3): ")
				fmt.Scanln(&choice)
				fmt.Println()

				switch choice { // these are the operations the Non-admin can perform
				case 1:
					viewEmployeeDetails(user)
				case 2:
					updateEmployeeDetails(user)
				case 3:
					fmt.Println("Logout successful!")

					return
				default:
					fmt.Println("Invalid choice! Please try again.")
				}

				fmt.Println()
			}
		}
	} else {
		fmt.Println("Login failed! Invalid username or password.") //it means that success is false.
	}
}
