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

// Empty structs are often used as a unary value. Sometimes, we don't care what is passed through a channel. We care when and if it is passed.
// We can block and wait until something is sent on a channel using the following syntax:
// <-ch
func waitForDBs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan
	}
}

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}

//-------------------------------------------------------------------------------------------------

//BUFFERED CHANNELS
// We can provide a buffer length as the second argument to make() to create a buffered channel:
// ch := make(chan int, 100)

func addEmailsToQueue(emails []string) chan string {
	emailChannel := make(chan string, len(emails))

	for _, email := range emails {
		emailChannel <- email
	}
	close(emailChannel)

	return emailChannel
}

//-------------------------------------------------------------------------------------------------

// CLOSING CHANNELS
func countReports(numSentCh chan int) int {
	total := 0
	for {
		numSentCh, ok := <-numSentCh
		if !ok {
			break
		}
		total += numSentCh
	}
	return total
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}

//-------------------------------------------------------------------------------------------------

// RANGE in CHANNELS
func concurrentFib(n int) {
	chInts := make(chan int)
	go func() {
		fibonacci(n, chInts)
	}()
	for v := range chInts {
		fmt.Println(v)
	}
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

//-------------------------------------------------------------------------------------------------

// SELECT
func logMessages(chEmails, chSms chan string) {
	for {
		select {
		case email, ok := <-chEmails:
			if !ok {
				return
			}
			logEmail(email)

		case sms, ok := <-chSms:
			if !ok {
				return
			}
			logSms(sms)
		}
	}
}

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmail(email string) {
	fmt.Println("Email:", email)
}

//-------------------------------------------------------------------------------------------------

//Select Default Case

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot(logChan)
		case <-saveAfter:
			saveSnapshot(logChan)
			return
		default:
			waitForData(logChan)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}

func main() {
	// test("Hello there Kaladin!")
	// test("Hi there Shallan!")
	// test("Hey there Dalinar!")

	//CHANNELS
	// emails := []email{
	// 	{"Email 1", time.Date(2019, time.December, 15, 0, 0, 0, 0, time.UTC)},
	// 	{"Email 2", time.Date(2021, time.January, 10, 0, 0, 0, 0, time.UTC)},
	// 	{"Email 3", time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)},
	// }
	// filterOldEmails(emails)

	//Select Default
	snapshotTicker := time.NewTicker(2 * time.Second).C
	saveAfter := time.After(10 * time.Second)
	logChan := make(chan string)

	go saveBackups(snapshotTicker, saveAfter, logChan)

	// Log messages from the channel
	for msg := range logChan {
		fmt.Println(msg)
	}
}
