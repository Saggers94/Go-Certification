package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	answer := 7 + 3*4
	fmt.Println("Answer: ", answer)

	answer = (7+3) * 4
	fmt.Println("Answer: ", answer)
	
	//multiplication
	radius := 12.0
	area := math.Pi * radius * radius 

	fmt.Println("Area is: ", area)

	// Integer Division works with only whole numbers
	half := 1 / 2
	fmt.Println("half with integer division", half)

	halfFloat := 1.0/2.0
	fmt.Println("half float:", halfFloat)

	//squaring(raising to the power) ('^' is a bitwise xor operation)
	badThreeSquared := 3^2
	fmt.Println("bad three squared", badThreeSquared)

	goodThreeSquared := math.Pow(3.0, 2.0)
	fmt.Println("Good Three Squared", goodThreeSquared)

	//modulus
	remainder := 50 % 3
	fmt.Println("remainder: ", remainder)

	// Unary Operator
	x := 3
	x++
	fmt.Println("x is now: ", x)

	x--
	x--
	fmt.Println("x is now: ", x)

	// Precedence
	// In multiplication and devidation, the parenthesis are redundunt
	a := 12.0 * 3.0 /4.0
	b := (12.0 * 3.0) / 4.0
	c := 12.0 * (3.0/4.0)

	fmt.Println("a", a, "b", b, "c", c)

	// Integer division
	unclear := 12 * (3/4)
	fmt.Println("unclear:", unclear)

	// Parenthesis
	f := 12.0/3.0/4.0
	fmt.Println("f", f)

	f = 12.0 / (3.0 / 4.0)
	fmt.Println("f is now", f)

	// Addition and Subtraction (Here parenthesis doesn't make any difference)
	fmt.Println()
	w := 12 + 3 - 4
	y := (12 + 3) - 4
	z := 12 + (3 - 4)
	fmt.Println("w", w, "y", y, "z", z)
	
	// precedence : parenthesis, multiplication/division, addition/subtraction
	w = 12 + 3 * 4
	y = (12 + 3) * 4
	z = 12 + (3 * 4)
	fmt.Println("w", w, "y", y, "z", z)

	// Modulus operator
	// Does one number devide exactly into another?
	p := 12
	q := 3
	if p%q == 0{
		fmt.Println(p, "divides exactly into", q)
	}else{
		fmt.Println(p, "does not divides exactly into", q)
	}

	// thisMonth := 4
	// fmt.Println("The month after", thisMonth, "is", thisMonth + 1)

	for m := 1; m<=12; m++ {
		fmt.Println("The month after", m, "is", m%12 + 1)
	}

	second := 31
	minute := 1

	// "&&" and "||" have same precedence
	if minute < 59 && second + 1 > 59{
		minute++
	}
	fmt.Println(minute)

	t := 12
	u := 6

	// if u != 0 && divideTwoNumbers(t,u)==2{
	// 	fmt.Println("Found 2")
	// }
	// v:=divideTwoNumbers(t,u)
	// if v == 2{
	// 	fmt.Println("We found a two")
	// }

	v, err := divideTwoNumbers(t,u)
	if err != nil {
		fmt.Println(err)
	}else{
		if v == 2{
			fmt.Println("We found 2")
		}
	}

	// Assignment operation
	l := 12
	l++
	fmt.Println("l ", l)
	l--
	fmt.Println("l is now", l)

	k := 10
	k *= 2
	fmt.Println("k is", k) 
	k /= 2
	fmt.Println("k is", k) 

	// The below will give us an error
	// z := y -= 8
}

func divideTwoNumbers(x, y int) (int, error){
	if y == 0{
		return 0, errors.New("cannot divide by 0")
	}
	return x/y, nil
}