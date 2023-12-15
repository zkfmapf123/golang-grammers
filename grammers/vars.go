package main

import "fmt"

// global variables
var (
	name             = "Foo"
	firstName string = "Foo"
	lastName  string // empty value
)

// global variables (const)
const (
	strictName      = "Foo"
	strictFirstName = "Foo"
	strictLastName  = ""
)

func main() {
	// local variables (var)
	version := 10          // int
	otherversion := "Bar"  // string
	anotherVersion := 10.1 // float32 or float64
	var typeVersion = 1
	var typeInt int

	// local variables (const)
	const name = "leedonggyu"
	const age = 30

	fmt.Println(version)
}
