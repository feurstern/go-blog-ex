package main

import "fmt"

func main() {
	firstName, _ := getNames()
	fmt.Println("Welcome to Textio,", firstName)

	y, z := subs()
	fmt.Println("xoxox", z, y)

}

// don't edit below this line
// setter
func subs() (int, int) {
	var x int = 24
	var y int = 32
	return x, y
}
func getNames() (string, string) {
	return "John", "Doe"
}
