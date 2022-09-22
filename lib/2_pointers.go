package lib

import "fmt"

func Pointers() {
	color := "Green"
	fmt.Println("Color is currently", color)

	changeUsingPointer(&color)
	fmt.Println("Color is currently", color)

}

func changeUsingPointer(s *string) {
	*s = "Red"
}
