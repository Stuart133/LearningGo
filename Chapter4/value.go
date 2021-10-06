package main

import "fmt"

type data struct {
	name string
	age  int
}

func (d data) displayName() {
	fmt.Println("My name is", d.name)
}

func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "Is age", d.age)
}

func main() {
	d := data{name: "Bill"}

	f1 := d.displayName
	f1()
	d.name = "Joan"
	f1()
}
