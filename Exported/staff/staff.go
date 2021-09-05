package staff

import "log"

// If you define functions or variables or even types with a "Capital" letter than that
// is being exported
var OverPaidLimit = 75000

// This begins with a lowercase letter, so that wouldn't be exported
var underPaidLimit = 20000

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) OverPaid() []Employee {
	var overpaid []Employee

	for _, x := range e.AllStaff {
		if x.Salary > OverPaidLimit {
			overpaid = append(overpaid, x)
		}
	}

	return overpaid
}

func (e *Office) UnderPaid() []Employee{
	var underpaid []Employee
	myFunction()
	for _, x := range e.AllStaff{
		if x.Salary < underPaidLimit {
			underpaid = append(underpaid, x)
		}
	}

	return underpaid
}

func (e *Office) notVisible(){
	log.Println("Hello World!")
}

// This below function doesn't have a receiver so we won't be able to call with
// the variable office
func myFunction(){
	log.Println("I am a function")
}

// Types, Variables, Constants and Functions that are created with the "Uppercase" letter
// are visible outside of the package and those that begin with the Lowercase letter
// are not visible outside of the package that they are created from 

// Composition  means embedding types into types