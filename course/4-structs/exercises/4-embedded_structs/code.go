package main

import "fmt"

type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}

type classKnight struct {
	player
	mainAttribute string
	baseattack    int
}

func showClassKnight(k classKnight) {
	fmt.Println("The nickname", k.nickname)
	fmt.Println("The username", k.username)
	fmt.Println("The base attack:", k.baseattack)
	fmt.Println("The level :", k.level)
	fmt.Println("The main attribute:", k.mainAttribute)
}

type player struct {
	username string
	password string
	nickname string
	level    int
}

type student struct {
	firstname string
	lastname  string
	age       int
	major     string
}

type classCourse struct {
	student
	courseName string
	duration   int
}

func studentCourse(c classCourse) {
	fmt.Println("Student name :", c.firstname+" "+c.lastname)
	fmt.Println("Student major :", c.major)
	fmt.Println("student age:", c.age)
	fmt.Println("student course class:", c.courseName)
}

// don't edit below this line

func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("====================================")
}

func main() {
	test(sender{
		rateLimit: 10000,
		user: user{
			name:   "Deborah",
			number: 18055558790,
		},
	})

	studentCourse(classCourse{
		student: student{
			firstname: "Hatsune",
			lastname:  "Miku",
			age:       17,
			major:     "Computer Science",
		},
		courseName: "Fullstack web dev",
		duration:   10,
	})

	showClassKnight(classKnight{
		player: player{
			username: "Hatsune Miku",
			password: "ulan bator",
			nickname: "hehehhe",
			level:    99,
		},
		mainAttribute: "magic",
		baseattack:    78,
	})

	test(sender{
		rateLimit: 5000,
		user: user{
			name:   "Sarah",
			number: 19055558790,
		},
	})
	test(sender{
		rateLimit: 1000,
		user: user{
			name:   "Sally",
			number: 19055558790,
		},
	})
}
