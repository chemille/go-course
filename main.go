package main

import (
	"fmt"
	"strconv"
)

const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50 //uint means it cannot be a negative
var bookings = make([]map[string]string, 0) // this creates an empty list of maps, with an initial size of 0 which can expand

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketCount := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketCount {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings: %v\n", firstNames)

			if remainingTickets == 0 {
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

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, element := range bookings { // underscore is a blank identifier to ignore a variable you don't want to use
		firstNames = append(firstNames, element["firstName"])
	}
	return firstNames
}

// moved validateUserInput func into helper.go

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) // & is a pointer that will print out the memory address of the variable

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Create an empty map for user
	// All keys have the same data type, and all values have the same data type
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	// FormatUint takes our uint value and formats it to a string as a decimal number
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
 	
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
