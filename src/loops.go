package main

import "fmt"

func Greet_Loop(str string, strs []string, do func(string), times int) {
	for i := 0; i < times; i++ {
		do("1-" + str)
	}

	// while loop
	j := 0
	for j < times {
		// break
		if j >= times {
			break
		}

		// continue
		if j%2 == 0 {
			j++
			continue
		}

		do("2-" + str)
		j++
	}

	// infinite loop
	/* for {
	   // ...
	 } */

	// range
	/* types of range
	Array or slice
	String
	Map
	Channel
	*/
	for index, value := range strs {
		fmt.Print(index)
		fmt.Println(": " + value)
	}
}

func main_loop() {
	var doFunc = func(str string) {
		fmt.Println(str)
	}

	slice := []string{
		"Zee",
		"MM",
		"Doe",
	}
	Greet_Loop("Zee", slice, doFunc, 5)
}
