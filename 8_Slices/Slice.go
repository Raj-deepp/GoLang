package main

import "fmt"

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}

	costs := [3]int{}
	cumulativeLength := 0

	for i, message := range messages {
		cumulativeLength += len(message)
		costs[i] = cumulativeLength
	}
	return messages, costs
}

func main() {
	messages, costs := getMessageWithRetries("hello", "world", "!")
	fmt.Println("Messages:", messages)
	fmt.Println("Costs:", costs)
}
