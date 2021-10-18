package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	waitForResult()
	fanOut()
}

func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "data"
		fmt.Println("Child : send signal")
	}()

	d := <-ch
	fmt.Println("parent : recv'd signal :", d)

	time.Sleep(time.Second)
	fmt.Println("-----------------------------------")
}

// Be careful with backpressure from this pattern - Think about the total goroutine count
func fanOut() {
	// Use a buffered channel to stop individual goroutines being blocked
	children := 2000
	ch := make(chan string, children)

	// Create a goroutine for each unit of work
	for c := 0; c < children; c++ {
		go func(child int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "data"
			fmt.Println("child : sent signal :", child)
		}(c)
	}

	// We use a counter to ensure all signals are received, rather than checking each individual signal
	for children > 0 {
		d := <-ch
		children--
		fmt.Println(d)
		fmt.Println("parent : recv'd signal :", children)
	}

	time.Sleep(time.Second)
	fmt.Println("-----------------------------------")
}
