package main

import (
	"fmt"
	"strconv"
)

//ERROR
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

//Formating Strings
func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs $%.2f to be sent to %v can not be sent", cost, recipient)
}


//The Error Interface
type divideError struct {
	dividend float64
}

func(de divideError) Error() string{
    return fmt.Sprintf("can not divide %v by zero", de.dividend)
}
//--------------------------------------------------------------------------------------------

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

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
