package lib

import (
	"fmt"
)

func Basics() {
	fmt.Println("Hello, world!")

	var whatToSay string = "Goodbye, world."
	fmt.Println(whatToSay)

	var number int = 4
	fmt.Println("Number is set to ", number, ". Cool!")

	whatWasSaid, anotherThing := saySomething() // Shorthand to create a variable without type and var prefix.
	fmt.Println("The function returned", whatWasSaid, anotherThing)

}

func saySomething() (string, string) {
	return "something", "another thing"
}
