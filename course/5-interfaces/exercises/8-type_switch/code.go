package main

import (
	"fmt"
)

func getExpenseReport(e expense) (string, float64) {
	switch c := e.(type) {
	case email:
		return c.toAddress, c.cost()
	case sms:
		return c.toPhoneNumber, c.cost()

	default:
		return "", 0.0

	}

}

type score struct {
	studentId, value1, value2 int
}

func (s score) average() float32 {
	return (float32(s.value1) + float32(s.value2)) / 2
}

type studentReport interface {
	average() float32
}

func getStudentReport(s studentReport) (string, float32) {

	switch c := s.(type) {
	case score:
		return "A", (float32(c.value1) + float32(c.value2)) / 2
	}

	return "b", 0.0
}

// don't touch below this line

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func estimateYearlyCost(e expense, averageMessagesPerYear int) float64 {
	return e.cost() * float64(averageMessagesPerYear)
}

func test(e expense) {
	address, cost := getExpenseReport(e)
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	case sms:
		fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	default:
		fmt.Println("Report: Invalid expense")
		fmt.Println("====================================")
	}
}

func showsStudentReport(s score) {
	fmt.Printf("The final score iss %f", s.average())
}

func main() {
	// test(email{
	// 	isSubscribed: true,
	// 	body:         "hello there",
	// 	toAddress:    "john@does.com",
	// })
	// test(email{
	// 	isSubscribed: false,
	// 	body:         "This meeting could have been an email",
	// 	toAddress:    "jane@doe.com",
	// })
	// test(email{
	// 	isSubscribed: false,
	// 	body:         "Wanna catch up later?",
	// 	toAddress:    "elon@doe.com",
	// })
	// test(sms{
	// 	isSubscribed:  false,
	// 	body:          "I'm a Nigerian prince, please send me your bank info so I can deposit $1000 dollars",
	// 	toPhoneNumber: "+155555509832",
	// })
	// test(sms{
	// 	isSubscribed:  false,
	// 	body:          "I don't need this",
	// 	toPhoneNumber: "+155555504444",
	// })
	// test(invalid{})

	student := score{
		value1:    90,
		value2:    78,
		studentId: 121211,
	}

	getStudentReport(student)

	showsStudentReport(student)
}
