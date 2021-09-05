package main

import "fmt"

func main() {
	apples := 18
	oranges := 9

	// Boolean Expression
	fmt.Println(apples == oranges)
	fmt.Println(apples != oranges)

	// >,<,>=,<=
	// You can only compare values that have the same types
	// if apples == "10"{

	// }

	fmt.Printf("%d > %d: %t", apples, oranges, apples > oranges)
	fmt.Println()

	fmt.Printf("%d > %d: %t", apples, oranges, apples >= oranges)
	fmt.Println()

	fmt.Printf("%d > %d: %t", apples, oranges, apples < oranges)
	fmt.Println()

	fmt.Printf("%d > %d: %t", apples, oranges, apples <= oranges)
	fmt.Println()
}