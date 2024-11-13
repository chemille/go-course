package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	// You can declare var like this too --> conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50 //uint means it cannot be a negative
	var bookings []string

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")

	// In Go, there are only for loops. No while loops.
	// infinite for loop
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// User input
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName) // & is a pointer that will print out the memory address of the variable

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName) // & is a pointer that will print out the memory address of the variable

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email) // & is a pointer that will print out the memory address of the variable

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets) // & is a pointer that will print out the memory address of the variable

		// User input validation
		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketCount := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketCount { // if all 3 are true
			remainingTickets = remainingTickets - userTickets
			// Dynamically adding to array
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, element := range bookings { // underscore is a blank identifier to ignore a variable you don't want to use
				var names = strings.Fields(element)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is fully booked. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entereddoes not contain @ sign")
			}
			if !isValidTicketCount {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}

}
