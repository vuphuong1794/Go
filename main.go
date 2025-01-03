package main

import "fmt"

func main() {
	fmt.Println("hello go")

	var name string
	var age int

	fmt.Scan(&name, &age)
	fmt.Printf("ten %v tuoi %v", name, age)
}
