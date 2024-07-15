package main

import "fmt"

// define interface
type Animal interface {
	Speak() string
}

// define structure
type Cat struct{}

func (c Cat) Speak() string {
	return "喵喵"
}

type Dog struct{}

func (d Dog) Speak() string {
	return "汪汪"
}

func main() {
	var animal Animal
	animal = Cat{}
	fmt.Println(animal.Speak())

	animal = Dog{}
	fmt.Println(animal.Speak())
}
