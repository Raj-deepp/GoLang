package main

import (
	"fmt"
	// "time"
)

func billingCost(plan string) float64 {
	switch plan {
	case "basic":
		return 10.0
	case "pro":
		return 20.0
	case "enterprise":
		return 50.0
	default:
		return 0.0
	}
}

func main() {
	// messageLen := 20
	// maxMessageLen := 50
	// fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	// if messageLen <= maxMessageLen {
	// 	fmt.Println("Message sent")
	// } else {
	// 	fmt.Println("Message not sent")
	// }

	// email:="abc@gmail.com"
	// if length := len(email); length > 10 {
    // fmt.Println("Email is valid")
	// }
	// fmt.Println("Email is invalid")

	//FUNCTION billingcost
	plan := "basic"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "pro"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "enterprise"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "free"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "unknown"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
}
