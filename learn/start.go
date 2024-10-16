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
func main() {
	messageStart:="Happy birthday! You are now"
	age:=21
	messageEnd:="years old!"

	fmt.Println(messageStart, age, messageEnd)
}
