package main

import "fmt"

// user defined type
type Salutation string
type SalutationObj struct {
	name     string
	greeting string
}

// constants
const (
	PI       = 3.14
	Language = "Go"
)
const (
	A = iota // succession in integer
	B
	C
)

func variable_main() {
	var message string
	var a, b, c int
	message = "hello, Go!!!"

	var message1 string = "hello, Go2!!!"
	var a1, b1, c1 int = 1, 2, 3

	var message2 = "hello, Go3!!!"
	var a2, b2, c2 = 1, 2, 3

	// can only use this short version inside a function
	message3 := "hello, Go4!!!"
	a3, b3, c3 := 1, false, 3

	// pointer: & is the memory address
	var greeting *string = &message

	fmt.Printf(message)
	fmt.Println(message, a, b, c)
	fmt.Println(message1, a1, b1, c1)
	fmt.Println(message2, a2, b2, c2)
	fmt.Println(message3, a3, b3, c3)

	fmt.Println(message, greeting, *greeting)
	*greeting = "hi, Zee"
	fmt.Println(message, greeting, *greeting)

	var message4 Salutation = "hello, Go5!!!"
	fmt.Println(message4)
	var s = SalutationObj{"Zee", "Hello, Go6!!!"}
	fmt.Println(s.name, s.greeting)
	var s1 = SalutationObj{name: "Zee1", greeting: "hello, Go7!!!"}
	fmt.Println(s1.name, s1.greeting)
	var s2 = SalutationObj{}
	s2.name = "Zee2"
	fmt.Println(s2.name, s2.greeting)

	fmt.Println(PI, Language, A, B, C)
}
