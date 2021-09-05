package main

import "fmt"

// If we define an interface that has two functions and than we define the functions
// with the Receiver that is related to type than by default this interface is extending those types
//and so, we can use the functionality of the interface
// Interface type
type Animal interface {
	Says() string
	HowManyLegs() int
}

func (d *Dog) Says() string{
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func (d *Cat) Says() string{
	return d.Sound
}

func (d *Cat) HowManyLegs() int {
	return d.NumberOfLegs
}

func main() {
	//ask a riddle
	dog := Dog {
		Name: "dog",
		Sound: "woof",
		NumberOfLegs: 4,
	}

	Riddle(&dog)

	var cat Cat
	cat.Name = "cat"
	cat.NumberOfLegs = 4
	cat.Sound ="Meow"
	cat.HasTail = true

	Riddle(&cat)
}

func Riddle(a Animal) {
	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs. what animal is it?`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}