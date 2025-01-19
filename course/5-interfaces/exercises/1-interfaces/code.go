package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) {
	fmt.Println("Th birth date :", msg.getMessage())
}

type message interface {
	getMessage() string
}

type mathOperation interface {
	addition() int
	// substraction() int
	// multiply() int
}

func result(c mathOperation) {
	fmt.Println(c.addition())
	// fmt.Println(c.substraction())
	// fmt.Println(c.multiply())
}

type mathOperator struct {
	x int
	y int
	z int
}

func (c mathOperator) addition() int {
	return c.x + c.y + c.z
}

func (c mathOperator) subsraction() int {
	return c.x - c.y - c.z
}

func (c mathOperator) multiply() int {
	return c.x * c.y * c.z
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func test(m message) {
	sendMessage(m)
	fmt.Println("====================================")
}

func main() {
	test(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})

	result(mathOperator{
		x: 10,
		z: 2,
		y: 3,
	})

	test(birthdayMessage{
		recipientName: "John Doe",
		birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
	// test(sendingReport{
	// 	reportName:    "First Report",
	// 	numberOfSends: 10,
	// })
	// test(birthdayMessage{
	// 	recipientName: "Bill Deer",
	// 	birthdayTime:  time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
	// })
}
