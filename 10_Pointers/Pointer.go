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

func main() {
	//Ref & Deref
	message := "This AB is a test message with CDE profanity."
	removeProfanity(&message)
	fmt.Println(message)
}
