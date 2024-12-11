package main

import (
	"fmt"
	"time"
)

// CONCURRENCY
// Concurrency is as simple as using the go keyword when calling a function.
func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}
func test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

//-------------------------------------------------------------------------------------------------

// CHANNELS
// Channels are a typed, thread-safe queue. Channels allow different goroutines to communicate with each other.
// ch := make(chan int)
// ch <- 69		//Send
// v := <-ch	//Receive
type email struct {
	body string
	date time.Time
}

func filterOldEmails(emails []email) {
	isOldChan := make(chan bool)

	// Correct time.Date and use goroutine properly
	go func() {
		for _, e := range emails {
			isOldChan <- e.date.Before(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC))
		}
		close(isOldChan)
	}()

	// Read results from channel
	i := 1
	for isOld := range isOldChan {
		fmt.Printf("email %d is old: %v\n", i, isOld)
		i++
	}
}

//-------------------------------------------------------------------------------------------------

func main() {
	// test("Hello there Kaladin!")
	// test("Hi there Shallan!")
	// test("Hey there Dalinar!")

	//CHANNELS
	emails := []email{
		{"Email 1", time.Date(2019, time.December, 15, 0, 0, 0, 0, time.UTC)},
		{"Email 2", time.Date(2021, time.January, 10, 0, 0, 0, 0, time.UTC)},
		{"Email 3", time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)},
	}
	filterOldEmails(emails)
}
