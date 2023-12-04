package main

import (
	"fmt"
	"os"
	"sync"
)

var correctRuns = 0

func onlyIncrements() {
	var wg sync.WaitGroup
	var locker sync.Mutex

	increments := 1000
	counter := 0

	increment := func() {
		defer wg.Done()

		locker.Lock()
		defer locker.Unlock()
		counter++
	}

	wg.Add(increments)
	for i := 0; i < increments; i++ {
		go increment()
	}

	wg.Wait()

	if counter == increments {
		correctRuns++
	} else {
		fmt.Fprintf(os.Stderr, `Calculations were not correct!
		Expected %d, got %d
		Correct runs: %d
		`, counter, increments, correctRuns)
		os.Exit(1)
	}
}

func main() {
	numRuns := 100

	for i := 0; i < numRuns; i++ {
		onlyIncrements()
	}
	fmt.Println("Calculations were correct!")
}
