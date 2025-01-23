package main

import (
	"fmt"
)

func getExpenseReport(e expense) (string, float64, string) {
	c, ok := e.(email)

	if ok {
		return c.toAddress, c.cost(), "nice"
	}

	s, ok := e.(sms)

	if ok {
		return s.toPhoneNumber, s.cost(), "nice"
	}
	return "", 0.0, "err"
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

func (u user) fullName() string {
	return u.firstName + " " + u.lastName
}

func userInfo(e showUser) (string, int) {
	c, ok := e.(user)

	if ok {
		return c.firstName + " " + c.lastName, c.age
	}
	return "", 0
}

type showUser interface {
	fullName() string
}

type user struct {
	firstName, lastName string
	age                 int
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

func displayUser(u user) {
	// fmt.Printf("FUllname : %s", u.fullName())
	fullName, age := userInfo(u)

	fmt.Printf("Fullname: %s Age : %d", fullName, age)
}

func test(e expense) {
	address, cost, t := getExpenseReport(e)
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	case sms:
		fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	default:
		fmt.Printf("Report: Invalid expense :%s \n", t)
		fmt.Println("====================================")
	}
}

func main() {

	e := email{
		isSubscribed: true,
		body:         "The sun will eventually die when it run out fuel as hidrogen",
		toAddress:    "naziincolour123@gmail.com",
	}

	test(e)
	test(email{
		isSubscribed: true,
		body:         "hello there",
		toAddress:    "john@does.com",
	})

	displayUser(user{firstName: "hatsune", lastName: "miku", age: 17})
	// test(email{
	// 	isSubscribed: false,
	// 	body:         "This meeting could have been an email",
	// 	toAddress:    "jane@doe.com",
	// })
	// test(email{
	// 	isSubscribed: false,
	// 	body:         "This meeting could have been an email",
	// 	toAddress:    "elon@doe.com",
	// })
	// test(sms{
	// 	isSubscribed:  false,
	// 	body:          "This meeting could have been an email",
	// 	toPhoneNumber: "+155555509832",
	// })
	// test(sms{
	// 	isSubscribed:  false,
	// 	body:          "This meeting could have been an email",
	// 	toPhoneNumber: "+155555504444",
	// })
	// test(invalid{})
}
