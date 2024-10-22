package main
import "fmt"

// func main(){
// 	fmt.Println("Hello World")
// }


// func main(){
// 	var username string="nowicnw"
// 	var password string="3298547698"
// 	fmt.Println("Authorization: Basic ", username+": "+password)
// }


//Hard way
// func main() {
// 	var smsSendingLimit int
// 	var costPerSMS float64
// 	var hasPermission bool
// 	var username string
	
// 	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
// }


//The WALRUS OPERATOR, := , declares a new variable and assigns a value to it in one line.
// func main() {
// 	messageStart:="Happy birthday! You are now"
// 	age:=21
// 	messageEnd:="years old!"

// 	fmt.Println(messageStart, age, messageEnd)
// }


// func main() {
// 	numMessagesFromDoris := 72
// 	costPerMessage := .02
// 	totalCost := costPerMessage * float64(numMessagesFromDoris)
// 	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
// }


//Same line Declaration
// func main() {
// 	averageOpenRate, displayMessage:= .23,"is the average open rate of your messages"
// 	fmt.Println(averageOpenRate, displayMessage)
// }


//Converting b/w datatypes
// func main() {
// 	accountAgeFloat := 2.6
// 	accountAgeInt:=int(accountAgeFloat)
// 	fmt.Println("Your account has existed for", accountAgeInt, "years")
// }


//Creating Constants
func main() {
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
}

