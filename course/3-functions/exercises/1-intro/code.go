package main

import "fmt"

func concat(s1, s2 string) string {
	return s1 + s2
}

func main() {
	fmt.Println(concat("Lane,", " happy birthday!"))
	test("Elon,", " hope that Tesla thing works out")
	test("Go", " is fantastic")
	result1 := multiplication(4, 25)

	fmt.Printf("The result is %d", result1)

}

func test(s1, s2 string) {
	fmt.Println(concat(s1, s2))
}
