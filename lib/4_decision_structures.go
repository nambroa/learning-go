package lib

import "fmt"

func DecisionStructures() {
	// If
	cat := "cat"

	if cat == "cat" && !false {
		fmt.Println("Cat is cat.")
	} else if false {
		fmt.Println("What are you doing here?")
	}

	// Switch
	animal := "dog"

	switch animal {
	case "cat":
		fmt.Println("Animal is cat")
	case "dog":
		fmt.Println("Animal is dog")
	}
}
