package main

import (
	"fmt"
)

type messageToSend struct {
	phoneNumber int
	message     string
}

func test(m messageToSend) {
	fmt.Println("Sending message to ", m.phoneNumber)
	fmt.Println("Message: ", m.message)
}

func main() {
	test(messageToSend{1234567890, "Hello World"})
	test(messageToSend{9876543210, "Hello phuong"})
}
