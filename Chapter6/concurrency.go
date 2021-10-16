package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		lowercase()
		wg.Done()
	}()

	go func() {
		uppercase()
		wg.Done()
	}()

	fmt.Printf("Waiting to finish\n")
	wg.Wait()

	fmt.Printf("\nTerminating program")
}

func lowercase() {
	fmt.Printf("lowercase yo")
}

func uppercase() {
	fmt.Printf("UPPERCASE YO")
}

func init() {
	runtime.GOMAXPROCS(1)
}