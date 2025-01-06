package main

import "fmt"

func main() {
	var firstName string
	var lastName string
	conferenceName := "GO Conference"
	const conferenceTickets = 50
	var remaingTickets uint = 50

	//var bookings [50]string //array (static size)
	var bookings2 []string //slice (dynamic size)
	var userTickets uint

	fmt.Println("Welcome to the " + conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remaingTickets)
	fmt.Println("Please enter your name to book a ticket")

	//for i := 0; i < 50; i++ {}

	for {
		fmt.Println("Enter your firstName and last name: ")
		fmt.Scan(&firstName, &lastName)

		fmt.Println("Enter number of tickets you want to book: ")
		fmt.Scan(&userTickets)

		if userTickets > remaingTickets {
			fmt.Printf("Sorry %v %v, we only have %v tickets remaining \n", firstName, lastName, remaingTickets)
			continue
		}

		remaingTickets = remaingTickets - userTickets
		fmt.Printf("Thank you %v %v for booking %v tickets. \n", firstName, lastName, userTickets)

		//bookings[0] = "booking 1"

		bookings2 = append(bookings2, firstName+" "+lastName)

		fmt.Printf("these are all our bookings: %v \n", bookings2)
		//fmt.Printf("ten %v tuoi %v", firstName, lastName)
	}
}
