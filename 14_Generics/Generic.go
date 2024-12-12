package main

import (
	"errors"
	"fmt"
)

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

// CONSTRAINTS
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

//-------------------------------------------------------------------------------------------------

// Parametric Constraints
type biller[C customer] interface {
	Charge(C) bill
	Name() string
}

// don't edit below this line

type userBiller struct {
	Plan string
}

func (ub userBiller) Charge(u user) bill {
	amount := 50.0
	if ub.Plan == "pro" {
		amount = 100.0
	}
	return bill{
		Customer: u,
		Amount:   amount,
	}
}

func (sb userBiller) Name() string {
	return fmt.Sprintf("%s user biller", sb.Plan)
}

type orgBiller struct {
	Plan string
}

func (ob orgBiller) Name() string {
	return fmt.Sprintf("%s org biller", ob.Plan)
}

func (ob orgBiller) Charge(o org) bill {
	amount := 2000.0
	if ob.Plan == "pro" {
		amount = 3000.0
	}
	return bill{
		Customer: o,
		Amount:   amount,
	}
}

type customer interface {
	GetBillingEmail() string
}

type bill struct {
	Customer customer
	Amount   float64
}

type user struct {
	UserEmail string
}

func (u user) GetBillingEmail() string {
	return u.UserEmail
}

type org struct {
	Admin user
	Name  string
}

func (o org) GetBillingEmail() string {
	return o.Admin.GetBillingEmail()
}

//-------------------------------------------------------------------------------------------------

func main() {
	// firstInts, secondInts := splitAnySlice([]int{0, 1, 2, 3})
	// fmt.Println(firstInts, secondInts)
}
