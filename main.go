package main

import "fmt"

func main() {
	var firstName string
	var lastName string
	var bookings [50]string //array (static size)
	var bookings2 []string  //slice (dynamic size)

	for {
		fmt.Println("Enter your firstName and last name: ")
		fmt.Scan(&firstName, &lastName)

		bookings[0] = "booking 1"

		bookings2 = append(bookings2, firstName+" "+lastName)

		fmt.Printf("these are all our bookings: %v \n", bookings2)
		//fmt.Printf("ten %v tuoi %v", firstName, lastName)
	}
}
