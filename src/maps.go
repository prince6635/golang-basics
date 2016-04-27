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

	// simple way to create a map
	prefixMap1 := map[string]string{
		"Bob":  "Mr ",
		"Joe":  "Dr ",
		"Amy":  "Dr ",
		"Mary": "Mrs ",
	}

	// update map, insert if key doesn't exist
	prefixMap1["Amy"] = "Jr "

	// delete in map
	// DON'T have to check whether the key exists or not.
	delete(prefixMap1, "Amy")

	// prefixMap1["Mary"] returns 2 values
	if value, exists := prefixMap1["Mary"]; exists {
		return "Found! " + value
	}

	return prefixMap[name] + ", " + prefixMap1["Amy"]
}

// Fields is the concrete LogFields type, as in myLogger.WithFields(bark.Fields{"foo": "bar"}).Info("Fields!")
type Fields map[string]interface{}

func PrintGeneralInterfaceValuesMap() {
	var data map[string]interface{}
	data = make(map[string]interface{})
	data["Zee"] = "Mr "
	data["Joe"] = 123
	fmt.Printf("%+v\n", data)

	fields := Fields{
		"Bob":  "Mr ",
		"Joe":  123,
		"Amy":  "Dr ",
		"Mary": 456,
	}
	fmt.Printf("%+v\n", fields)

	for key, val := range fields {
		fmt.Printf("key:%v, value:%v\n", key, val)
	}
}

func main() {
	fmt.Println(GetPrefix("Zee"))
	PrintGeneralInterfaceValuesMap()
}
