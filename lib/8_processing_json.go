package lib

import (
	"encoding/json"
	"fmt"
)

type SuperHero struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func ProcessingJson() {
	myJson := `[
	{
		"first_name": "Clark",
		"last_name": "Kent",
		"hair_color": "black",
		"has_dog": true
		
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne",
		"hair_color": "black",
		"has_dog": false
		
	}
]`
	var unmarshalled []SuperHero // Processing from JSON to "DTO" is called unmarshalling.

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		fmt.Println("Error unmarshalling json", err)
	}

	fmt.Printf("unmarshalled: %v", unmarshalled) // Json has been written to a struct

	// Writing from a Struct TO a JSON
	var mySlice []SuperHero
	var marshallOne SuperHero
	marshallOne.FirstName = "Wally"
	marshallOne.LastName = "West"
	marshallOne.HairColor = "red"
	marshallOne.HasDog = false

	var marshallTwo SuperHero
	marshallTwo.FirstName = "Diana"
	marshallTwo.LastName = "Prince"
	marshallTwo.HairColor = "black"
	marshallTwo.HasDog = true

	mySlice = append(mySlice, marshallOne, marshallTwo)

	newJson, err := json.MarshalIndent(mySlice, "", "   ")
	if err != nil {
		fmt.Println("error marhsalling", err)
	}

	fmt.Println(string(newJson)) // Converts json from bytes to string.

}
