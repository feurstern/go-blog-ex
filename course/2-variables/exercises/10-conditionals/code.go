package main

import "fmt"

func main() {

	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	var priceText = 0.2
	var strMessage string = "Theere were a ghost in theee towns who haunts the peopel here"

	totalMsgLength := len(strMessage)
	totalPay := priceText * float64(totalMsgLength)

	fmt.Printf("total pay %.2f", totalPay)
	fmt.Println(strMessage)

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}

	// messageLen > 0 ? fmt.Println("xoxoxo") : fmt.Println("xixixix")
}
