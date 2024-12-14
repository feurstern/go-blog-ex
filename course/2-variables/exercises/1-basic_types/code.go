package main

import "fmt"

func main() {
	// initialize variables here
	var smsSendingLimit int = 23
	var costPerSMS float64 = 4.5
	var hasPermission bool = true
	var username string = "Hatsunee"

	var a int = 24
	var b int = 5
	var c int = b * a

	fmt.Print("The result of %v * %v = %v", a, b, c)

	fmt.Printf(" The messsage limit :%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
