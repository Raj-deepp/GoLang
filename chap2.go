package main

import (
	"fmt"
	// "time"
)

func main() {
	// the current time can only be known when the program is running
	// attempts to assign the result of time.Now() to a constant, which is not allowed in Go. Constants in Go are evaluated at compile-time.
	// const currentTime = time.Now()
	
	//Corrrect Way
	// currentTime := time.Now()
	// fmt.Println("Current time: ", currentTime)

	// s := fmt.Sprintf("I am %v years old", "way too many")
	// fmt.Println(s)


	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent\n",name,openRate)
	fmt.Print(msg)
}
