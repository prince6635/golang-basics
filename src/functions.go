package main

import "fmt"

type SalutationForFunc struct {
	name     string
	greeting string
}

func Greet(salute SalutationForFunc) {
	// fmt.Println(salute.name, salute.greeting)
	fmt.Println(CreateMessage(salute.name, salute.greeting))

	message, alternate := CreateMessageWithAlternate(salute.name, salute.greeting)
	fmt.Println(message)
	fmt.Println(alternate)

	message1, _ := CreateMessageWithAlternate("Zee", "Hello again!!!")
	fmt.Println(message1)
	// can't declare but don't use it
	// fmt.Println(alternate1)

	message2, alternate2 := CreateMessageWithUncertainParams("Zee", "hello1", "hello2")
	fmt.Println(message2)
	fmt.Println(alternate2)
}

func CreateMessage(name, greeting string) string {
	return greeting + " " + name
}

// return multiple values
func CreateMessageWithAlternate(name, greeting string) (string, string) {
	return greeting + " " + name, "HEY! " + name
}

func CreateMessageWithDeclaredReturnValues(name, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "HEY! " + name
	return
}

// variadic function
func CreateMessageWithUncertainParams(name string, greetings ...string) (message string, alternate string) {
	fmt.Println(len(greetings))
	message = greetings[0] + " " + name
	alternate = "HEY! " + greetings[1] + " " + name
	return
}

// function types
// func GreetByFuncType(salute SalutationForFunc, printFunc func(string)) {

// Better to define a function type
type Printer func(string) // () means no return value
func GreetByFuncType(salute SalutationForFunc, printFunc Printer) {
	message, alternate := CreateMessageWithDeclaredReturnValues(salute.name, salute.greeting)
	printFunc(message)
	printFunc(alternate)
}

func Print(str string) {
	fmt.Print(str)
}

func PrintLine(str string) {
	fmt.Println(str)
}

// closure
func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(custom + ", " + s)
	}
}

func main_functions() {
	var s = SalutationForFunc{"Zee", "Hello, function!!!"}
	Greet(s)

	GreetByFuncType(s, Print)
	GreetByFuncType(s, PrintLine)

	GreetByFuncType(s, CreatePrintFunction("Custom by closure:"))
}
