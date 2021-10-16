package main

import (
	"crypto/sha1"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		printHashes("A")
		wg.Done()
	}()

	go func() {
		printHashes("B")
		wg.Done()
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

func printHashes(prefix string) {
	for i := 1; i < 50000; i++ {
		num := strconv.Itoa(i)
		sum := sha1.Sum([]byte(num))
		fmt.Printf("%s: %05d: %x\n", prefix, i, sum)
	}

	fmt.Println("Completed", prefix)
}

func init() {
	runtime.GOMAXPROCS(1)
}
