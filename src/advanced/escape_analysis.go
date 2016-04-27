package main

import "fmt"

type Point struct {
	x, y int
}

// the compile see you return the address and just makes it on the heap for you.
// This is a common idiom in go.
// otherwise, will return by value, http://stackoverflow.com/questions/10866195/stack-vs-heap-allocation-of-structs-in-go-and-how-they-relate-to-garbage-collec
func newPoint() *Point {
	return &Point{10, 20}
}

// Escape Analysis in Go
// https://scvalex.net/posts/29/
var intPointerGlobal *int = nil

func Foo() *int {
	anInt0 := 0
	anInt1 := new(int)
	*anInt1 = 1

	anInt2 := 42
	intPointerGlobal = &anInt2

	anInt3 := 5

	fmt.Printf("inside Foo - 0:%v, 1:%v, 2:%v, 3:%v, intPointerGobal:%v\n",
		anInt0, *anInt1, anInt2, anInt3, *intPointerGlobal)
	return &anInt3
}

func main() {
	fmt.Printf("inside main - Foo:%v\n", *Foo())
	fmt.Printf("inside main - newPoint:%+v\n", *newPoint())
}
