package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstName string
	var lastName string
	conferenceName := "GO Conference"
	const conferenceTickets = 50
	var remaingTickets uint = 50

	//var bookings [50]string //array (static size)
	var bookings []string //slice (dynamic size)
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

		bookings = append(bookings, firstName+" "+lastName)

		firstNames := []string{}
		for _, booking := range bookings { //(index, booking: mỗi phần tử trong bookings sẽ có giá trị riêng biệt, sử dụng _ để không bị lỗi khôg dùng index)
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("these are all our bookings: %v \n", bookings)
		fmt.Printf("these are all our first names: %v  \n", firstNames)
	}
}
