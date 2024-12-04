package main

import (
	"errors"
	// "fmt"
)

//Maps
func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap:=make(map[string]user)
	if len(names) != len(phoneNumbers){
		return nil, errors.New("Invalid Sizes")
	}
	for i:=0; i<len(names);i++{
		name:=names[i]
		phoneNumber:=phoneNumbers[i]
		userMap[name]=user{
			name:name,
			phoneNumber:phoneNumber,
		}
	}
	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}//------------------------------------------------------------------------------------------------



func main(){

}