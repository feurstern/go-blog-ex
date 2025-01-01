package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// ?
func (c authenticationInfo) getBasicAuth() string {
	return c.username + ":" + c.password
}

type rectangle struct {
	length int
	width  int
}

func (x rectangle) recCircumference() int {
	return 2 * x.length * x.width
}

func result(x rectangle) {
	fmt.Println(x.recCircumference())

}

const currentYear = 2024

type user struct {
	firstname string
	lastname  string
	yearBorn  int
}

func (c user) userInfo() string {
	return c.firstname + " " + c.lastname
}

func showUserInfo(c user) {
	fmt.Println("Full name", c.userInfo())
	age := currentYear - c.yearBorn
	fmt.Printf("age %d", age)
}

// don't touch below this line

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("====================================")
}

func main() {

	result(rectangle{
		width:  10,
		length: 5,
	})

	showUserInfo(user{
		firstname: "Hatsune",
		lastname:  "Miku",
		yearBorn:  2008,
	})
	// test(authenticationInfo{
	// 	username: "Google",
	// 	password: "12345",
	// })
	// test(authenticationInfo{
	// 	username: "Bing",
	// 	password: "98765",
	// })
	// test(authenticationInfo{
	// 	username: "DDG",
	// 	password: "76921",
	// })
}
