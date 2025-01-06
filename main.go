package main

import "fmt"

func main() {
	fmt.Println("hello go")

	var firstName string
	var lastName string
	var bookings [50]string //array
	var bookings2 []string  //slice

	fmt.Println("Enter your name and last name: ")
	fmt.Scan(&firstName, &lastName)

	bookings[0] = "booking 1"

	bookings2 = append(bookings2, firstName+" "+lastName)

	fmt.Printf("The whole array: %v \n", bookings)
	fmt.Printf("The first element: %v \n", bookings[0])
	fmt.Printf("Array type: %T \n", bookings)
	fmt.Printf("Array length: %v \n\n", len(bookings))

	fmt.Printf("All bookings in slice: %v \n", bookings2)
	//fmt.Printf("ten %v tuoi %v", firstName, lastName)
}
