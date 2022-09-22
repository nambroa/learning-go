package lib

import "fmt"

// Contents of the Interface details what must be satisfied by the types wanting to use it.
type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

func (d Dog) Says() string {
	return "Woof"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

type Gorilla struct {
	Name  string
	Color string
}

func (g Gorilla) Says() string {
	return "Kekw"
}

func (g Gorilla) NumberOfLegs() int {
	return 2
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}
func Interfaces() {
	dog := Dog{Name: "Samson", Breed: "German Shepherd"}

	PrintInfo(&dog)

	gorilla := Gorilla{Name: "GigaChad", Color: "Missigno"}

	PrintInfo(&gorilla)

}
