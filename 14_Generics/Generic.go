package main

import "errors"

// Put simply, generics allow us to use variables to refer to specific types. This is an amazing feature because it allows us to write abstract functions that drastically reduce code duplication.

// func splitAnySlice[T any](s []T) ([]T, []T) {
//     mid := len(s)/2
//     return s[:mid], s[mid:]
// }

// In the example above, T is the name of the type parameter for the splitAnySlice function, and we've said that it must match the any constraint, which means it can be anything. This makes sense because the body of the function doesn't care about the types of things stored in the slice.

func getLast[T any](s []T) T {
	if len(s) == 0 {
		var zeroVal T
		return zeroVal
	}
	return s[len(s)-1]
}

//CONSTRAINTS
type lineItem interface {
	GetCost() float64
	GetName() string
}

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	newBalance := balance - newItem.GetCost()
	if newBalance < 0.0 {
		return nil, 0.0, errors.New("Insufficient Funds")
	}
	oldItems = append(oldItems, newItem)
	return oldItems, newBalance, nil
}

func main() {
	// firstInts, secondInts := splitAnySlice([]int{0, 1, 2, 3})
	// fmt.Println(firstInts, secondInts)
}
