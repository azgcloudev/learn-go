package main

import (
	"fmt"
	"strings"
	"booking-app/helper"
)

var conferenceName = "Go conference"

const conferenceTickets = 50

var remainingTickets uint = 50

// var bookings [50]string // array declaration; fixed size
// slice
// var bookingSlice []string
// var bookingsSlice = string[]  {}
var bookingSlice = []string{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, email, uint(userTickets), remainingTickets)

		if isValidEmail && isValidName && isValidUserTickets {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Printf("Out conference is booked out. Come back next year.")
				break
			} else {
				if !isValidName {
					fmt.Printf("First name or last name you entered is too short.\n")
				}
				if !isValidEmail {
					fmt.Printf("email address you entered doesn't contain @\n")
				}
				if !isValidUserTickets {
					fmt.Printf("number of tickets you entered is invalid\n")
				}
			}
		}
	}

	// var lastName string

	/*
		fmt.Printf("\nThe whole array: %v\n", bookings)
		fmt.Printf("The whole slice: %v\n", bookingSlice)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Array type: %T\n", bookings)
		fmt.Printf("Slice type: %T\n", bookingSlice)
		fmt.Printf("Array length: %v\n", len(bookings))
		fmt.Printf("Slice length: %v\n", len(bookingSlice))
	*/

	// for loop
	for i := 0; i < 2; i++ {
		// do somthing 2 times
	}

	// infinite loop
	// for {

	// 	city := "london"

	// 	switch city {
	// 	case "New York":
	// 		// excecute code
	// 	case "London", "Berling", "Singapore":
	// 	// execute code
	// 	case "Mexico":
	// 	//execute code
	// 	default:
	// 		fmt.Printf("No valid city")

	// 	}

	// 	if isValidEmail && isValidName && isValidUserTickets {
	// 		if userTickets > int(remainingTickets) {
	// 			fmt.Printf("%v tickets exceeded the remaining available tickets: %v\n", userTickets, remainingTickets)
	// 			continue
	// 		}

	// 		bookings[0] = firstName + " " + lastName
	// 		bookingSlice = append(bookingSlice, firstName+" "+lastName)

	// 		remainingTickets = remainingTickets - uint(userTickets)

	// 		fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s\n", firstName, lastName, userTickets, email)

	// 		// for each loops
	// 		// _ =  black identifies;
	// 		// for index, booking := range bookingSlice
	// 		for _, booking := range bookingSlice {
	// 			if remainingTickets == 0 {
	// 				// end program
	// 				fmt.Println("Our conference is booked out. Come back next year.")
	// 				break
	// 			} else {

	// 				var names = strings.Fields(booking)
	// 				firstNames = append(firstNames, names[0])

	// 				fmt.Printf("The first names of bookings are:\n")
	// 				for _, name := range firstNames {
	// 					fmt.Printf("%v\n", name)
	// 				}
	// 				fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	// 			}

	// 		}
	// 	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here o attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookingSlice {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}


func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter the last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("How many tickets you want to buy? ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookingSlice = append(bookingSlice, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
