package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (person *Person) Rename(newName string) {
	person.name = newName
}

// interface, no need to be implemented,
// any named type that implements the methods in the interface can be used in the same way.
type Renamable interface {
	Rename(newName string)
}

func RenameToFrog(concretObj Renamable) {
	concretObj.Rename("Frog")
}

// Empty interface: can be used for any type
func TypeSwitchWithEmptyInterface(xType interface{}) {
	switch xType.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Person:
		fmt.Println("Person")
	default:
		fmt.Println("Unknown")
	}
}

// Flexibility of using interface in Go
/* https://golang.org/pkg/io/
type Writer

type Writer interface {
        Write(p []byte) (n int, err error)
}
Writer is the interface that wraps the basic Write method.
*/
// Now we implement this Writer interface in named type: Person
func (person *Person) Write(p []byte) (n int, err error) {
	str := string(p)
	person.Rename(str)
	n = len(str)
	err = nil
	return
}

func main_interfaces() {
	person := Person{"Zi", 100}
	fmt.Printf("%v\n", person.name)
	// RenameToFrog(person) // pass by pointer, need to deallocate
	RenameToFrog(&person)
	fmt.Printf("%v\n", person.name)

	// empty interface
	TypeSwitchWithEmptyInterface(1)
	TypeSwitchWithEmptyInterface("1")
	TypeSwitchWithEmptyInterface(person)

	// fmt.Fprintf(w io.Writer, format string, a ...interface{})
	// now we can use my named type: Person as the Writer interface and pass it to Fprintf function
	fmt.Fprintf(&person, "This is a counter %d", 10)
	fmt.Printf("%v\n", person.name)
}
