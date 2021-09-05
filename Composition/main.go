package main

import "fmt"

type Vehicle struct {
	NumberOfWheels     int
	NumberOfPassangers int
}

type Car struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehicle    Vehicle
}

// This function is tied with the Vehicle struct
func (v Vehicle) showDetails() {
	fmt.Println("Number of Passangers:", v.NumberOfPassangers)
	fmt.Println("Number of Wheels:", v.NumberOfWheels)
}

// This function is tied with the Car struct
func (c Car) show(){
	fmt.Println("Make:", c.Make)
	fmt.Println("Model:", c.Model)
	fmt.Println("Year:", c.Year)
	fmt.Println("Is Electric", c.IsElectric)
	fmt.Println("Is Hybrid", c.IsHybrid)
	c.Vehicle.showDetails();
}

// Using Composition instead of inheritance

func main() {
	suv := Vehicle{
		NumberOfWheels: 4,
		NumberOfPassangers: 6,
	}

	volvoXC90 := Car {
		Make: "Volvo",
		Model: "XC90 T8",
		Year: 2021,
		IsElectric: false,
		IsHybrid: true,
		Vehicle: suv,
	}

	volvoXC90.show()

	teslaModelx := Car{
		Make: "Tesla",
		Model: "Model x",
		Year: 2021,
		IsElectric: true,
		IsHybrid: false,
		Vehicle: suv,
	}

	// The nice thing about this is unlike OOP, we do not need getters and setters
	teslaModelx.Vehicle.NumberOfPassangers = 7

	teslaModelx.show()
	
}