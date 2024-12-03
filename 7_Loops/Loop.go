package main

import (
	"fmt"
)

//Format for loop:
// for INITIAL; CONDITION; AFTER{
//   // do something
// }

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
	}
	return totalCost
}

//---------------------------------------------------------------------------------------

// Omitting conditions from a for loop in Go
func maxMessages(thresh int) int {
	Tcost := 0
	for i := 0; ; i++ {
		Tcost += 100 + (1 * i)
		if Tcost > thresh {
			return i
		}
	}
}

//-------------------------------------------------------------------------------------------

//There is no WHILE loop in Go:
// for CONDITION {
//   // do some stuff while CONDITION is true
// }

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(maxCostInPennies) - actualCostInPennies
	for balance > 0 {
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
	}
	if balance < 0 {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}

//--------------------------------------------------------------------------------------------

// Fizzbuzz
func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizznuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	// numMessages := 10
	// totalCost := bulkSend(numMessages)
	// fmt.Printf("The total cost for sending %d messages is: $%.2f\n", numMessages, totalCost)

	// threshold := 1000
	// maxMsgs := maxMessages(threshold)
	// fmt.Printf("The maximum number of messages that can be sent without exceeding %d is: %d\n", threshold, maxMsgs)

	//Type of While loop in Go
	plantHeight := 1
	for plantHeight < 5 {
		fmt.Println("still growing! current height:", plantHeight)
		plantHeight++
	}
	fmt.Println("plant has grown to ", plantHeight, "inches")

	//Fizzbuzz
	fizzbuzz()
}
