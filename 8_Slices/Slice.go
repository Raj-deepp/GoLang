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

//Variadic Operator

func concat(strs ...string) string {
	final := ""
	// strs is just a slice of strings
	for i := 0; i < len(strs); i++ {
		final += strs[i]
	}
	return final
}

func sum(nums ...int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total = total + nums[i]
	}
	return total
}

//-------------------------------------------------------------------------------------------------

//Spread Operator

func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

//-------------------------------------------------------------------------------------------------

//Append

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	costByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costByDay) {
			costByDay = append(costByDay, 0.0)
		}
		costByDay[cost.day] += cost.value
	}
	return costByDay
}
//-------------------------------------------------------------------------------------------------

func main() {
	messages, costs := get1MessageWithRetries("hello", "world", "!")
	fmt.Println("Messages:", messages)
	fmt.Println("Costs:", costs)

	//concat variadic
	final := concat("Hello ", "there ", "friend!")
	fmt.Println(final)

	sumOfNums := sum(1, 2, 3, 4)
	fmt.Println(sumOfNums)

	//Spread (names...)
	names := []string{"bob", "sue", "alice"}
	printStrings(names...)
}
