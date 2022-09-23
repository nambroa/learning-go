package lib

import (
	"fmt"
	"math/rand"
	"time"
)

func insertNewNumber(intChan chan int) {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Int()
	intChan <- randomNumber
}

func Channels() {
	intChan := make(chan int)
	defer close(intChan) // Closes the channel when it has finished running through this function.

	go insertNewNumber(intChan) // Run this function in parallel.

	num := <-intChan // Listen for new values on this channel and assign it to num.
	fmt.Println(num)
}
