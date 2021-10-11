package main

import "fmt"

type Dog struct {
	Name       string
	IsMammal   bool
	PackFactor int
}

func (d *Dog) Speak() {
	fmt.Println("Woof!", "My name is", d.Name, ", it is", d.IsMammal, "I am a mammal with a pack factor of", d.PackFactor)
}

func (d *Dog) Update(i int) {
	d.PackFactor = i
}

type Cat struct {
	Name        string
	IsMammal    bool
	ClimbFactor int
}

func (c *Cat) Speak() {
	fmt.Println("Meow!", "My name is", c.Name, ", it is", c.IsMammal, "I am a mammal with a climb factor of", c.ClimbFactor)
}

func (c *Cat) Update(i int) {
	c.ClimbFactor = i
}

type Speaker interface {
	Speak()
	Update(i int)
}

func main() {
	speakers := []Speaker{
		&Dog{
			Name:       "Fido",
			IsMammal:   true,
			PackFactor: 5,
		},
		&Cat{
			Name:        "Milo",
			IsMammal:    true,
			ClimbFactor: 4,
		},
	}

	for _, speaker := range speakers {
		speaker.Update(2)
	}

	for _, speaker := range speakers {
		speaker.Speak()
	}
}
