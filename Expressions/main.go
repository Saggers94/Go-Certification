package main

import "fmt"

func main() {
	// These are assigned int,string,boolean literal
	age := 10
	name := "Jack"
	rightHanded := true

	fmt.Printf("%s is %d years old. Right handed:%t", name, age, rightHanded)
	fmt.Println()

	ageInTenYears := age + 10

	fmt.Printf("In ten years, %s will be %d years old", name, ageInTenYears)
	fmt.Println()

	isATeenager := age >= 13
	fmt.Println(name, "is a teenager", isATeenager)

	// Expression is a bit of code that can be evaluated to a single value
}