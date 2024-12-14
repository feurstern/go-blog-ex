package main

import (
	"fmt"
)

func yearsUntilEvents(age int) (int, int, int) {
	var yearsUntilAdult int = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	var yearsUntilDrinking int = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	var yearsUntilCarRental int = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return yearsUntilAdult, yearsUntilCarRental, yearsUntilDrinking
}

func triangle(side int) (int, int, int) {
	var length int = 4
	var height int = 5

	var result int = side * length * height

	return length, height, result
}

// don't edit below this line
func test(age int) {
	fmt.Println("Age:", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(age)
	fmt.Println("You are an adult in", yearsUntilAdult, "years")
	fmt.Println("You can drink in", yearsUntilDrinking, "years")
	fmt.Println("You can rent a car in", yearsUntilCarRental, "years")
	fmt.Println("====================================")
}

func main() {
	side := 4
	l, h, result := triangle(side)
	fmt.Println("l", l)
	fmt.Printf("h", h)
	fmt.Printf("the result is %d", result)
	test(4)
	test(10)
	test(22)
	test(35)
}
