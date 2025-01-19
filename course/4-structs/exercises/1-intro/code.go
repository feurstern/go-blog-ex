package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

type skills struct {
	programingLanguage string
	language           string
}

type Candidates struct {
	skill          skills
	age            int
	education      string
	isHasReference bool
}

type userAccount struct {
	username     string
	userPassword string
	userEmail    string
	firstName    string
	age          int
	isVip        bool
}

// don't edit below this line
func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("====================================")
}

func candidateMssg(c Candidates) {
	fmt.Printf("Candidate information  language : %s and %s ", c.skill.programingLanguage, c.skill.language)
}

func main() {

	candidate := Candidates{skill: skills{language: "c", programingLanguage: "cobol"}, age: 20, education: "bachelor", isHasReference: true}

	candidateMssg(candidate)
	// user := userAccount{firstName: "Hatsune Miku", isVip: true}
	// b := ""
	// if !user.isVip {
	// 	b = "You got free ticket to Manila"
	// } else {
	// 	b = "You got free ticker to Gulag"
	// }

	// fmt.Printf("Hi %s! Welcome to Gulag! %s\n", user, b)
	// test(messageToSend{
	// 	phoneNumber: 148255510981,
	// 	message:     "Thanks for signing up",
	// })
	// test(messageToSend{
	// 	phoneNumber: 148255510982,
	// 	message:     "Love to have you aboard!",
	// })
	// test(messageToSend{
	// 	phoneNumber: 148255510983,
	// 	message:     "We're so excited to have you",
	// })
}
