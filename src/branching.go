package main

import (
	"fmt"

	"./lib"
)

func main_branching() {
	// var s = lib.Salutation_Greeting{"Zi", "Hello"}
	// var s = lib.Salutation_Greeting{"Doe", "Hello"}
	var s = lib.Salutation_Greeting{"0123456789", "Hello"}
	lib.Greet(s, lib.CreatePrinterFunction_Greeting("customized"), false)
	fmt.Println("............")
	lib.Greet(s, lib.CreatePrinterFunction_Greeting("customized"), true)

	lib.TypeSwitchTest(1)
	lib.TypeSwitchTest(s)
	lib.TypeSwitchTest("string")
	lib.TypeSwitchTest(1.123)
}
