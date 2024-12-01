package main

import (
	"fmt"
	"strconv"
)

//ERROR Interface
func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	costForCustomer, err :=sendSMS(msgToCustomer)
    if err != nil{
        return 0.0, err
    }
    costForSpouse, err :=sendSMS(msgToSpouse)
    if err != nil{
        return 0.0, err
    }
    return costForCustomer + costForSpouse, nil
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
//-----------------------------------------------------------------------------------------

func main() {
	// Atoi converts a stringified number to an integer
	i, err := strconv.Atoi("42b")
	if err != nil {
		fmt.Println("couldn't convert:", err)
		// Output:
		// couldn't convert: strconv.Atoi: parsing "42b": invalid syntax
		return
	}
	// If we get here, the variable i was converted successfully
	fmt.Println("Converted integer:", i)
}
