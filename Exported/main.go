package main

import (
	"log"
	"myapp/staff"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 30000, FullTime: true},
	{FirstName: "Sally", LastName: "Jones", Salary: 50000, FullTime: true},
	{FirstName: "Mark", LastName: "Smithers", Salary: 60000, FullTime: true},
	{FirstName: "Allan", LastName: "Anderson", Salary: 15000, FullTime: false},
	{FirstName: "Margeret", LastName: "Carter", Salary: 100000, FullTime: false},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}


	// myStaff.All()
	// log.Println(myStaff.All())
	// staff.OverPaidLimit = 30000

	log.Println("OverPaid Staff",myStaff.OverPaid())
	log.Println("UnderPaid Staff",myStaff.UnderPaid())

	// We cannot use this as it begins with the small letter
	// myStaff.notVisible()
}