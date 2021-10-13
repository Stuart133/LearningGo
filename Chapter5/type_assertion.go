package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Car struct {
}

func (c Car) String() string {
	return "Car"
}

type Cloud struct {
}

func (c Cloud) String() string {
	return "Cloud"
}

func main() {
	rand.Seed(time.Now().UnixNano())

	mvs := []fmt.Stringer{
		Car{},
		Cloud{},
	}

	for i := 0; i < 10; i++ {
		rn := rand.Intn(2)

		if v, is := mvs[rn].(Cloud); is {
			fmt.Println("Got lucky:", v)
			continue
		}

		fmt.Println("Got unlucky")
	}
}
