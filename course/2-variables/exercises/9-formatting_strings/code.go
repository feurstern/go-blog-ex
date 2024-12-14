package main

import "fmt"

func main() {
	const name = "Hatsune Miku"
	const openRate = 30.5

	// ?
	// don't edit below this line
	const hourlyRate = 2.5
	const payRate = 90
	const jobsCount = 6
	totalPay := float64(jobsCount) * hourlyRate * payRate

	salarySummary := fmt.Sprintf(
		"Hi! Your today job count is %d. So the total is (%d * %.2f) * %.1f = %.2f",
		jobsCount, jobsCount, hourlyRate, payRate, totalPay,
	)
	msg := fmt.Sprintf("Thee messagee %s and  the open rate os %.1f", name, openRate)

	fmt.Println(msg)
	fmt.Println(salarySummary)
}
