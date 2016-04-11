package main

import "fmt"

/* Map:
keys have to have equality operator defined
maps are reference types (like a pointer)
not thread safe
*/

func GetPrefix(name string) (prefix string) {
	var prefixMap map[string]string
	prefixMap = make(map[string]string)

	prefixMap["Zee"] = "Mr "
	prefixMap["Joe"] = "Dr "
	prefixMap["Doe"] = "Dr "
	prefixMap["Mary"] = "Mrs "

	return prefixMap[name]
}

func main() {
	fmt.Println(GetPrefix("Zee"))
}
