package main

import (
	"fmt"
)

type user struct {
	name   string
	number int
}

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

func canSendMessage(msg messageToSend) bool {
	if msg.sender.name == "" {
		return false
	}
	if msg.recipient.name == "" {
		return false
	}
	if msg.recipient.number == 0 {
		return false
	}
	return true
}

func sendMessage(msg messageToSend) {
	if canSendMessage(msg) {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}

func main() {
	var sender, recipient user
	var message string

	// Nhập thông tin người gửi
	fmt.Println("Nhập thông tin người gửi:")
	fmt.Print("Tên: ")
	fmt.Scanln(&sender.name)
	fmt.Print("Số: ")
	fmt.Scanln(&sender.number)

	// Nhập thông tin người nhận
	fmt.Println("\nNhập thông tin người nhận:")
	fmt.Print("Tên: ")
	fmt.Scanln(&recipient.name)
	fmt.Print("Số: ")
	fmt.Scanln(&recipient.number)

	// Nhập nội dung tin nhắn
	fmt.Print("\nNhập nội dung tin nhắn: ")
	fmt.Scanln(&message)

	// Tạo message
	msg := messageToSend{
		message:   message,
		sender:    sender,
		recipient: recipient,
	}

	// Gửi tin nhắn
	sendMessage(msg)
}
