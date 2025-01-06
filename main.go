package main

import "fmt"

func main() {
	fmt.Println("hello go")

	// var name string
	// var age int
	var bookings [50]string

	bookings[0] = "booking 1"

	fmt.Printf("The whole array: %v \n", bookings)
	fmt.Printf("The first element: %v \n", bookings[0])
	fmt.Printf("Array type: %T \n", bookings)
	fmt.Printf("Array length: %v \n", len(bookings))

	//fmt.Scan(&name, &age)
	//fmt.Printf("ten %v tuoi %v", name, age)
}
