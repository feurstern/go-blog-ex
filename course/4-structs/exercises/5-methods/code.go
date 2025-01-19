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

type coordinatePos struct {
	x int
	y int
	z int
}

func (p coordinatePos) currentPosition() int {
	return 2 * (p.x + p.y + p.z)
}

type rectangle struct {
	length int
	width  int
}
type triangle struct {
	high int
	base int
}

func (x rectangle) recCircumference() int {
	return 2 * x.length * x.width
}

func (x triangle) recCircumference() int {
	return x.high * x.base
}

type shape interface {
	recCircumference() int
}

func result(x shape) {
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

	coordinate := coordinatePos{x: 10, y: 2, z: 5}

	println(coordinate.currentPosition())

	// result(rectangle{
	// 	width:  10,
	// 	length: 5,
	// })

	// result(triangle{
	// 	base: 10,
	// 	high: 5,
	// })

	// showUserInfo(user{
	// 	firstname: "Hatsune",
	// 	lastname:  "Miku",
	// 	yearBorn:  2008,
	// })
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
