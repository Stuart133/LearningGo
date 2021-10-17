package main

import (
	"fmt"
	"runtime"
	"sync"
)

func concurrent() {
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

	fmt.Println("\nTerminating program")
}

func lowercase() {
	fmt.Printf("lowercase yo")
	fmt.Printf("lowercase yo")
}

func uppercase() {
	fmt.Printf("UPPERCASE YO")
	fmt.Printf("UPPERCASE YO")
}

func init() {
	runtime.GOMAXPROCS(1)
}
