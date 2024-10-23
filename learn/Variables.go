package main
import (
	"fmt"
	// "time"
	// "unicode/utf8"
)

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
// func main() {
// 	const premiumPlanName = "Premium Plan"
// 	const basicPlanName = "Basic Plan"

// 	fmt.Println("plan:", premiumPlanName)
// 	fmt.Println("plan:", basicPlanName)
// }

func main() {
	// the current time can only be known when the program is running
	// attempts to assign the result of time.Now() to a constant, which is not allowed in Go. Constants in Go are evaluated at compile-time.
	// const currentTime = time.Now()
	
	//Corrrect Way
	// currentTime := time.Now()
	// fmt.Println("Current time: ", currentTime)

	// s := fmt.Sprintf("I am %v years old", "way too many")
	// fmt.Println(s)


	// const name = "Saul Goodman"
	// const openRate = 30.5

	// msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent\n",name,openRate)
	// fmt.Print(msg)


	//Go also has a special type, "rune", which is an alias for int32. This means that a rune is a 32-bit integer, which is large enough to hold any Unicode code point.
	// const name = "boots"
	// fmt.Printf("constant 'name' byte length: %d\n", len(name))
	// fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name))
	// fmt.Println("=====================================")
	// fmt.Printf("Hi %s, so good to have you back in the arcanum\n", name)

	// const name = "🐻" //The emoji is only one rune, but it takes up 4 bytes.
	// fmt.Printf("constant 'name' byte length: %d\n", len(name))
	// fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name))
	// fmt.Println("=====================================")
	// fmt.Printf("Hi %s, so good to have you back in the arcanum\n", name)


	fname := "Dalinar"
	lname := "Kholin"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."

	userLog:=fmt.Sprintf("Name: %s %s, Age: %d, Rate: %0.1f, Is Subscribed: %t, Message: %s",fname,lname,age,messageRate,isSubscribed,message)

	fmt.Println(userLog)
}
