package lib

import "fmt"

type Salutation_Greeting struct {
	// need to be capital so branching.go can construct this struct
	Name     string
	Greeting string
}

type Printer_Greeting func(string)

// if
func Greet(salutation Salutation_Greeting, do Printer_Greeting, isFormal bool) {
	msg, alternate := CreateMessage_Greeting(salutation.Name, salutation.Greeting)
	// if isFormal {
	// do(msg)

	// embedded if, just declare prefix, that's only used in if statement
	if prefix := getPrefix(salutation.Name); isFormal {
		do(prefix + msg)
	} else {
		do(alternate)
	}
}

// switch
// !!! default will break in each case except explicitly telling it to fall through
func getPrefix(name string) (prefix string) {
	switch name {
	case "Zi":
		prefix = "Mr "
	case "M":
		prefix = "Mrs "
	case "Doe":
		prefix = "Dr "
	default:
		prefix = "Dude "
	}

	// more complicated switch
	switch name {
	case "Zi":
		fmt.Println("Zee")
		// if name is "Zi", it'll always fall through the next case "M"
		fallthrough
	case "M":
		fmt.Println("MM")
	case "Doe", "Joe":
		fmt.Println("Doe or Joe")
	default:
		fmt.Println("Default")
	}

	// switch on nothing
	switch {
	case name == "Zi":
		fmt.Println("Zee")
	case name == "Joe", name == "Doe", len(name) == 10:
		fmt.Println("JD10")
	default:
		fmt.Println("Default")
	}

	return
}

// switch on type
// !!! here, interface{} means x can be any type,
// similar as Object in Java and void pointer in C++
func TypeSwitchTest(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Salutation_Greeting:
		fmt.Println("Salutation_Greeting")
	default:
		fmt.Println("unknown")
	}
}

func CreateMessage_Greeting(name, greeting string) (msg string, alternate string) {
	msg = greeting + " " + name
	alternate = "HEY! " + name
	return
}

func CreatePrinterFunction_Greeting(custom string) Printer_Greeting {
	return func(s string) {
		fmt.Println(custom + ": " + s)
	}
}
