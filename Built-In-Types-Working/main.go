package main

import "fmt"

//basic types (numbers,string,boolean)

// Always use int type
// var myInt int

// will never use this
// var myInt16 int16
// var myInt32 int32
// var myInt64 int64

// This is unsigned Int which can hold only positive value
// var myUint uint

// float64 can hold much larger number than float32
// var myFloat float32
// var myFloat64 float64

//aggregate types (struct, array)
// Arrays
// var myStrings [3]string
//Struct
//  type Car struct {
// 	NumberOfTires int
// 	Luxury bool
// 	BucketSeats bool
// 	Make string
// 	Model string
// 	Year int
//  }

type Animal struct{
	Name string
	Sound string
	NumberOfLegs int
}
//reference types (pointers, slices, maps, functions, channels)

//interface type


func main() {
	// myInt = 60
	// the below would give an error as the uint can only hold the positive value
	// myUint = -10
	// myUint = 10
	// myFloat = 10.1
	// myFloat64 = 100.1

	// log.Println(myInt,myUint, myFloat, myFloat64)

	// String in Go are immutable, so it could run in an performance issues
	// myString := "Sagar"
	// log.Println(myString)
	// myString = "John"

	// myBool := true
	// myBool = false
	// log.Println(myBool)


	//Aggregate types
	
	// var myStrings [3]int
	// myStrings[0] = "Cat"
	// myStrings[1] = "Dog"
	// myStrings[2] = "Fish"

 	// fmt.Println("First element in the array is", myStrings[0])

	//  One way for the struct to work
	//  var myCar Car
	//  myCar.NumberOfTires = 4
	//  myCar.BucketSeats = true
	//  myCar.Luxury = true
	//  myCar.Model = "Honda"
	//  myCar.Make = "Volkswagon"

	// Second way of the struct
	//  myCar := Car{
	// 	 NumberOfTires: 4,
	// 	 Luxury: true,
	// 	 BucketSeats: true,
	// 	 Make: "Volkswagon",
	// 	 Model: "Honda",
	// 	 Year: 1987,
	//  }
	//  fmt.Println("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)


	//Pointers

	// x := 10

	// This "&" before the x would give the address of the x in the memory
	// It points to the location in the memory
	// myFirstPointer := &x

	// fmt.Println("x is", x)
	// fmt.Println("My First Pointer is:", myFirstPointer)


	// "*" before the "myFirstPointer" would give the value that is stored in the address
	// *myFirstPointer = 15
	// fmt.Println("X is now", x)

	// What this has allowed is the changing of the variable value that is not in the 
	// scope of the function
	// changeValueOfPointer(&x)

	// fmt.Println("After function call, X is now", x)
	// var myInt int
	// myInt = 10

	// ftm.Println(myInt)


	// Slices 
	// var animals []string
	// animals = append(animals, "Dog")
	// animals = append(animals, "Fish")
	// animals = append(animals, "Cat")
	// animals = append(animals, "Horse")

	// fmt.Println(animals)

	// for i, x := range animals{
	// 	fmt.Println(i, x)
	// }

	// fmt.Println("Element 0 is", animals[0])
	
	// fmt.Println("First two elements are", animals[0:2])

	// fmt.Println("The slice is", len(animals))

	// fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	// sort.Strings(animals)

	// fmt.Println("Is it sorted now?", sort.StringsAreSorted(animals))

	// fmt.Println(animals)

	// animals = DeleteFromSlice(animals, 1)

	// fmt.Println(animals)

	//Map
	// You cannot create map like below
	// var myMap map[string] string

	// Map is pass by reference, so you don't need to worry about pointers
	// intMap := make(map[string]int)
	// intMap["one"] = 1
	// intMap["two"] = 2
	// intMap["three"] = 3
	// intMap["four"] = 4
	// intMap["five"] = 5

	// Maps are not sorted and cannot be sorted
	// for key, value := range intMap {
	// 	fmt.Println(key, value)
	// }

	// deleting in Map. Map are very very fast
	// delete(intMap, "four")

	// el, ok := intMap["four"]

	// if ok {
	// 	fmt.Println(el, "is in Map.")
	// }else{
	// 	fmt.Println(el, "is not in Map.")
	// }

	// intMap["two"] = 4


	// Functions
	// z := addTwoNumber(1,3)
	// fmt.Println(z)

	myTotal := sumMany(2,3,4,5,6,88,7,-4,-5)
	fmt.Println(myTotal)

	var dog Animal
	dog.Name = "dog"
	dog.Sound = "Woof"
	dog.NumberOfLegs = 4

	cat := Animal {
		Name: "Cat",
		Sound: "Meow",
		NumberOfLegs: 4,
	}
	dog.Says()
	cat.Says()
	cat.howManyLegs()
}

// func changeValueOfPointer(num *int){
// 	*num = 25
// }

// func DeleteFromSlice(a []string, i int) []string{
// 	a[i] = a[len(a)-1]
// 	a[len(a)-1] = ""
// 	a = a[:len(a)-1]
// 	return a
// }

// (sum int) would name the return type
// func addTwoNumber(x, y int) int{
// 	return x+y
// }

// This is called variadic function that take "n" numbers of parameter
// "...int" means take any number of integer
func sumMany(nums ...int) int{
	total := 0

	for _, x := range nums{
		total = total + x
	}

	return total
}

// here the (a *Animal) is called as a receiver
// Receiver on a function to link a function to a type
func (a *Animal) Says(){
	fmt.Printf("A %s says %s", a.Name, a.Sound)
	fmt.Println()
}

func (a *Animal) howManyLegs(){
	fmt.Printf("A %s has %d legs", a.Name, a.NumberOfLegs)
}