package lib

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
}

// This function is tied to the person struct, you can call it directly through a person.
func (person *Person) printFirstName() string { // Input param is called a receiver.
	return person.FirstName
}

func TypesAndStructs() {
	john := Person{
		FirstName: "John",
		LastName:  "Doe",
	}
	fmt.Println(john.printFirstName())

	// Creates a map that maps string to People. Maps are NOT sorted by insertion.
	people := make(map[string]Person)
	people["john"] = john

	// Slices are Go's version of arrays (kinda).
	var numbers []int

	numbers = append(numbers, 2)
	numbers = append(numbers, 1)
	numbers = append(numbers, 3)
	// this can all be shorthanded to numbers := []int{2,1,3}

	sort.Ints(numbers)

	fmt.Println("Numbers are", numbers)
	fmt.Println("First number is", numbers[0])
}
