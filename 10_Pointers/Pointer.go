package main

import (
	"fmt"
	"strings"
)

// Reference & Dereference(*)
func removeProfanity(message *string) {
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "AB", "**")
	messageVal = strings.ReplaceAll(messageVal, "CDE", "***")
	*message = messageVal
}

//-------------------------------------------------------------------------------------------------

// NIL Pointers
func removeProfanityNil(message *string) {
	if message == nil {
		return
	}
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "AB", "**")
	messageVal = strings.ReplaceAll(messageVal, "CDE", "***")
	*message = messageVal
}

//-------------------------------------------------------------------------------------------------

// Pointer Receiver Code
func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func main() {
	//Ref & Deref
	message := "This AB is a test message with CDE profanity."
	removeProfanity(&message)
	fmt.Println(message)

	//Pointer Receiver Code
	myEmail := email{
		message:     "Hello, how are you?",
		fromAddress: "sender@example.com",
		toAddress:   "receiver@example.com",
	}
	fmt.Println("Before update:", myEmail.message)
	myEmail.setMessage("This is an updated message.")
	fmt.Println("After update:", myEmail.message)
}
