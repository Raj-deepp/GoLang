package main

import "fmt"

// func concat(s1 string, s2 string) string {     //FUNCTION SIGNATURE
//OR
func concat(s1, s2 string) string {
	return s1 + s2
}

//FUN 2
func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	lastMonthBill := getBillForMonth(costPerSend, numLastMonth)
	thisMonthBill := getBillForMonth(costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	return costPerSend * messagesSent
}


//FUN 3
//Ignoring Return Values by a "_"
func getProductMessage(tier string) string {
	quantityMsg, priceMsg, _ := getProductInfo(tier)
	return "You get " + quantityMsg + " for " + priceMsg + "."
}

func getProductInfo(tier string) (string, string, string) {
	if tier == "basic" {
		return "1,000 texts per month", "$30 per month", "most popular"
	} else if tier == "premium" {
		return "50,000 texts per month", "$60 per month", "best value"
	} else if tier == "enterprise" {
		return "unlimited texts per month", "$100 per month", "customizable"
	} else {
		return "", "", ""
	}
}


//Named Return Values
func getCoords() (x, y int){
  // x and y are initialized with zero values

  return // automatically returns x and y  Naked return
}


func yearsUntilEvents(age int) (yearsUntilAdult int, yearsUntilDrinking int, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return
}


//FUN 4
func reformat(message string, formatter func(string) string) string {
    formattedMessage := formatter(formatter(formatter(message)))
    return "TEXTIO: " + formattedMessage
}


//Anonymous Functions
func printReports(intro, body, outro string) {
	introCost := func(msg string) int {
		return 2 * len(msg)
	}
	
	bodyCost := func(msg string) int {
		return 3 * len(msg)
	}

	outroCost := func(msg string) int {
		return 4 * len(msg)
	}

	printCostReport(introCost, intro)
	printCostReport(bodyCost, body)
	printCostReport(outroCost, outro)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}


func main() {
	fmt.Println(concat("Lane,", " happy birthday!"))
	fmt.Println(concat("Elon,", " hope that Tesla thing works out"))
	fmt.Println(concat("Go", " is fantastic\n"))


	//FUN 2
	costPerSend := 5
	numLastMonth := 1000
	numThisMonth := 1500

	increase := monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth)
	fmt.Printf("The bill increased by %d pennies.\n\n", increase)


	//FUN 3
	tier := "premium" // You can change this to "basic", "premium", or "enterprise"
	message := getProductMessage(tier)
	fmt.Println(message)


	//FUN 4
	message4 := "General Kenobi"
    
    // Example formatter function that adds a period at the end of the string
    formatter := func(s string) string {
        return s + "."
    }
    
    // Use the reformat function with the message and formatter
    result := reformat(message4, formatter)
    
    // Print the result
    fmt.Println(result) // Output: TEXTIO: General Kenobi...
}
