package main

import (
	"errors"
	"fmt"
)

func get1MessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}

	costs := [3]int{}
	cumulativeLength := 0

	for i, message := range messages {
		cumulativeLength += len(message)
		costs[i] = cumulativeLength
	}
	return messages, costs
}

//Slices in Go
// primes := [6]int{2, 3, 5, 7, 11, 13}
// mySlice := primes[1:4]
// // mySlice = {3, 5, 7}

//SYNTAX
// arrayname[lowIndex:highIndex]
// arrayname[lowIndex:]
// arrayname[:highIndex]
// arrayname[:]

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()
	if plan == "planpro" {
		return allMessages[:], nil
	}
	if plan == "planfree" {
		return allMessages[0:2], nil
	}
	return nil, errors.New("Unsupported Plan")
}

func getMessageWithRetries() [3]string {
	return [3]string{
		"Click",
		"Here",
		"Please",
	}
}

//-------------------------------------------------------------------------------------------------

//MAKE

// func make([]T, len, cap) []T
// mySlice := make([]int, 5, 10)

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		cost := float64(len(message)) * 0.01
		costs[i] = cost
	}
	return costs
}

//-------------------------------------------------------------------------------------------------

func main() {
	messages, costs := get1MessageWithRetries("hello", "world", "!")
	fmt.Println("Messages:", messages)
	fmt.Println("Costs:", costs)
}
