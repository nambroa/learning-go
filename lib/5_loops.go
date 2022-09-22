package lib

import "fmt"

func Loops() {
	// for loop - while or do while loops are NOT in go.
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	animals := []string{"dog", "fish", "horse", "cow", "cat"}
	// i refers to the index of the current animal being looped. You can change i to _ in case you don't need it.
	for i, animal := range animals {
		fmt.Println(i, animal)
	}

	otherAnimals := make(map[string]string)
	otherAnimals["dog"] = "Fido"
	otherAnimals["cat"] = "Fluffy"
	for animalType, animal := range animals {
		fmt.Println(animalType, animal)
	}
}
