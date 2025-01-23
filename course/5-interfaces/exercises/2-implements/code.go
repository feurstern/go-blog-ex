package main

import (
	"fmt"
)

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

// ?

type axis2d struct {
	x int
	y int
}

type axis3d struct {
	x int
	y int
	z int
}

type axisDiagonal struct {
	xNegative int
	xPostive  int
	yNegative int
	yPositive int
}

func (c axis2d) movement() int {
	return c.x + c.y
}

func (c axis3d) movement() int {
	return c.x + c.y + c.z
}

func (c axisDiagonal) movement() int {
	return (c.xNegative + c.xPostive) - (c.yNegative + c.yNegative)
}

type position interface {
	movement() int
}

func showCurrentPosition(c position) {
	fmt.Println("Current position", c.movement())
}

// don't touch below this line

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func test(e employee) {
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("====================================")
}

type knight struct {
	baseAttack  int
	baseHp      int
	baseMana    int
	baseArmor   int
	SwordDamage int
}

type archer struct {
	baseAttack int
	baseHp     int
	baseMana   int
	baseArmort int
	bowDamage  int
}

type playerClass interface {
}

func (c knight) playerStast() int {
	return c.baseAttack * c.SwordDamage
}

func (c archer) playerStats() int {
	return c.bowDamage * c.baseAttack
}

func showPlayerStatsInfo() {

}

func main() {
	test(fullTime{
		name:   "Jack",
		salary: 50000,
	})

	// showCurrentPosition(axis2d{
	// 	x: 4,
	// 	y: 10,
	// })

	// showCurrentPosition(axis3d{
	// 	x: rand.Intn(100),
	// 	y: rand.Intn(100),
	// 	z: rand.Intn(100),
	// })

	// showCurrentPosition(axisDiagonal{
	// 	xNegative: rand.Intn(200),
	// 	xPostive:  rand.Intn(200),
	// 	yNegative: rand.Intn(100),
	// 	yPositive: rand.Intn(200),
	// })

	test(contractor{
		name:         "Bob",
		hourlyPay:    100,
		hoursPerYear: 73,
	})
	test(contractor{
		name:         "Jill",
		hourlyPay:    872,
		hoursPerYear: 982,
	})
}
