package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func login() (User, bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n') // taking the input from the user
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n') // taking the input from the user
	password = strings.TrimSpace(password)

	for _, employee := range employees {
		if employee.Username == username && employee.Password == password { // if this condition is valid then you are successfully logged in
			return employee, true
		}
	}
	return nil, false
}
