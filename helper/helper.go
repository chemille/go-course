package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketCount := userTickets > 0 && userTickets <= remainingTickets
	// In Go, you can return multiple values
	return isValidName, isValidEmail, isValidTicketCount
}

// Need to explicitly export this helper func to be available for all packages in the app
// To export, you simply capitalize the first letter of the func
// You can also export: variables, constants, types 