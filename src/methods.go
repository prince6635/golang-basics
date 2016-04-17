package main

import (
	"fmt"
	// "./lib"
)

type NewSalutation struct {
	name     string
	greeting string
}

// create a method on a named type: steps 1, 2, 3
// step 1: named type
type NewSalutations []NewSalutation

type PrintFunc func(string, string)

func printGreetingMsg(name string, greeting string) {
	fmt.Println(name + ": " + greeting)
}

func GreetInFunction(salutations []NewSalutation, printFunc PrintFunc) {
	for _, s := range salutations {
		printFunc(s.name, s.greeting)
	}
}

// step 2: method in named type
func (salutations NewSalutations) GreetInNamedTypeMethod(printFunc PrintFunc) {
	for _, s := range salutations {
		printFunc(s.name, s.greeting)
	}
}

// remember to use a pointer if you need to modify the original object
func (newSalutation NewSalutation) Rename(newName string) {
	newSalutation.name = newName
}

func (newSalutation *NewSalutation) RenameByRef(newName string) {
	newSalutation.name = newName
}

func main_methods() {
	// salutation := lib.Salutation_Greeting{"Z", "Welcome, Z!"}
	salutation := NewSalutation{"Z", "Welcome, Z!"}
	fmt.Printf("%v: %v\n", salutation.name, salutation.greeting)
	salutation.Rename("Zee1")
	fmt.Printf("%v\n", salutation.name) // the original obj hasn't been changed
	salutation.RenameByRef("Zee2")
	fmt.Printf("%v\n", salutation.name)

	fmt.Println(".............")
	salutations := []NewSalutation{
		{"Z", "Welcome, Z!"},
		{"Zee", "Welcome, Zee!"},
	}
	GreetInFunction(salutations, printGreetingMsg)

	fmt.Println(".............")
	newSalutations := NewSalutations{
		{"Z", "Welcome, Z!"},
		{"Zee", "Welcome, Zee!"},
	}
	// step 3: use the method in named type
	newSalutations.GreetInNamedTypeMethod(printGreetingMsg)
}
