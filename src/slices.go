package main

import "fmt"

/*
Array:
  Fixed size, can't never be changed;
  Array type is decided by size and undlying type
  No initialization (all 0 valued)
  Not a pointer! - value type

Slice: a wrapper of an array, take part of the array, so the underlying has to be an array!
  Fixed size (but can be reallocated with append)
  Type is slice of underlying type (doesn't have length)
  Use make to initialize otherwise is nil
  Points to an array

!!! in function params, always pass slice instead of array
http://openmymind.net/The-Minimum-You-Need-To-Know-About-Arrays-And-Slices-In-Go/
In go, everything is pass by value.
However, a string is a struct that has a length and a pointer to a byte array.
When you pass a string to another function, it copies the length and the pointer.
As a consequence, the new copied string points to the same underlying data.
So pass []string instead of []*string, but for user-defined structs, better pass by []*MyClass,
https://www.goinggo.net/2013/09/iterating-over-slices-in-go.html
*/

type SliceObj struct {
	key string
	val string
}

func main_slices() {
	var s1 []int
	s1 = make([]int, 3) // make([]int, 3, 20) means capacity is 20
	s1[0] = 1
	s1[1] = 2
	s1[2] = 3
	// s1[3] = 4 // runtime error
	fmt.Printf("%v\n", s1)

	s2 := []int{1, 2, 3, 4}
	fmt.Printf("%v\n", s2)

	s3 := []SliceObj{
		{"key1", "Zee"},
		{"key2", "Z"},
		{"key3", "Zi"},
	}
	fmt.Printf("%v\n", s3)

	// slicing
	s4 := s3[1:2] // index starts from 1 and end before 2, = [1,2)
	s5 := s3[:2]  // = [0, 2)
	s6 := s3[1:]  // = [1, len(s3))
	fmt.Printf("%v\n", s4)
	fmt.Printf("%v\n", s5)
	fmt.Printf("%v\n", s6)

	// append
	s6 = append(s6, SliceObj{"key4", "NewZ"})
	fmt.Printf("%v\n", s6)
	s6 = append(s6, s6...) // double the slices, ... means all the elements in s6
	fmt.Printf("%v\n", s6)

	// delete: use appending 2 slices to remove the item
	s3 = append(s3[:1], s3[2:]...)
	fmt.Printf("%v\n", s3)

}
