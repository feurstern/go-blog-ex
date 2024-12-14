package main

import "fmt"

func main() {
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"
	// don't edit below this line
	userMessage, sessionTime, currentTime := "Hatsune Miku", 3, 24

	fmt.Println("Usermessage", userMessage, sessionTime, currentTime)
	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
}
