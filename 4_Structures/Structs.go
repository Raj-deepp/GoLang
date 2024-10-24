package main
import "fmt"

type car struct {
	brand      string
	model      string
	doors      int
	mileage    int
	frontWheel wheel
	backWheel  wheel
}

type wheel struct{
	radius int
	material string
}
//---------------------------------------------------------------

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" || mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.name == "" || mToSend.recipient.number == 0 {
		return false
	}
	return true
}
//--------------------------------------------------------------

//Embedded Structs
type car2 struct {
  make string
  model string
}

type truck struct {
  // "car" is embedded, so the definition of a
  // "truck" now also additionally contains all
  // of the fields of the car struct
  car2
  bedSize int
}
//------------------------------------------------------------------

func main(){
	// myCar := car{}
	// myCar.brand = "BMW"
	// myCar.frontWheel.radius = 5
	// fmt.Println(myCar)


	sender := user{name: "Alice", number: 12345}
	recipient := user{name: "Bob", number: 67890}

	msg := messageToSend{
		message:   "Hello!",
		sender:    sender,
		recipient: recipient,
	}

	canSend := canSendMessage(msg)
	fmt.Println("Can send message:", canSend)


	myCar2 := car2{}
	myTruck := truck{}
	myCar2.make = "TATA Motors"
	myCar2.model = "PUNCH"
	myTruck.make= "TATA Motors"

	fmt.Println("Car:", myCar2)
	fmt.Println("Truck:", myTruck)

}