package main

import (
	"fmt"
)

//Format for loop:
// for INITIAL; CONDITION; AFTER{
//   // do something
// }

func bulkSend(numMessages int) float64 {
	totalCost:=0.0
	for i:=0;i<numMessages;i++{
		totalCost += 1.0 + (0.01 * float64(i))
	}
	return totalCost
}
//---------------------------------------------------------------------------------------

//Omitting conditions from a for loop in Go
func maxMessages(thresh int) int {
	Tcost:=0
	for i:=0; ;i++{
		Tcost += 100 + (1 * i)
		if Tcost>thresh{
			return i
		}
	}
}


func main(){
	// numMessages := 10
	// totalCost := bulkSend(numMessages)
	// fmt.Printf("The total cost for sending %d messages is: $%.2f\n", numMessages, totalCost)

	threshold := 1000
	maxMsgs := maxMessages(threshold)
	fmt.Printf("The maximum number of messages that can be sent without exceeding %d is: %d\n", threshold, maxMsgs)
}