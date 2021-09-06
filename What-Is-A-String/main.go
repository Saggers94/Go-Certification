package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
	name := "Hello World!"
	fmt.Println("String:", name)
	fmt.Println()

	fmt.Println("Bytes")
	for i :=0 ; i<len(name); i++{
		fmt.Printf("%x ",name[i])
	}
	fmt.Println()
	fmt.Println()

	// Rune is a 32 bit integer
	fmt.Println("Index\tRune\tString")
	for x, y:= range(name){
		fmt.Println(x, "\t", y, "\t", string(y))
	}
	fmt.Println()

	name = "Γειά σου Κόσμε"
	fmt.Println("Index\tRune\tString")
	for x, y:= range(name){
		fmt.Println(x, "\t", y, "\t", string(y))
	}
	fmt.Println()

	// Join the string
	// String are immutable
	fmt.Println()
	fmt.Println("Three ways to concatenate strings")

	h := "Hello, "
	w := "World."

	// using + => Not very efficient
	// This is not the better way for concatinating
	myString := h+w
	fmt.Println(myString)

	// with "Sprintf", it is more efficient than by using the "+" sign
	// Using fmt => more efficient
	fmt.Println()
	fmt.Println(myString)
	myString = fmt.Sprintf("%s%s",h,w)

	//using "stringbuilder" => is very efficient and gives us more control
	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	fmt.Println(sb.String())

	// Substring
	fmt.Println()
	name = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println("Getting a substring")
	// How to get subset of the slices
	fmt.Println(name[10:13])


	// Index start at 0
	courseName := "Learn Go for Beginners Crash Course"

	// If you don't use the ":" in courseName[0:12], than it would print out the Rune
	fmt.Println(courseName[0])
	fmt.Println(string(courseName[0]))

	fmt.Println(string(courseName[6]))

	for i:=0;i<=21;i++{
		fmt.Print(string(courseName[i]))
	}

	fmt.Println()
	for i:=13;i<=21;i++{
		fmt.Print(string(courseName[i]))
	}


	// string length
	fmt.Println("Length of coursename is:", len(courseName))

	var mySlice []string
	mySlice = append(mySlice, "one")
	mySlice = append(mySlice, "two")
	mySlice = append(mySlice, "three")

	fmt.Println("mySlice has", len(mySlice), "elements")
	fmt.Println("The last element in my Slice is", mySlice[len(mySlice) - 1])

	courses := []string{
		"Learn Go for Beginners Crash Course",
		"Learn Java for Beginners Crash Course",
		"Learn Python for Beginners Crash Course",
		"Learn C for Beginners Crash Course",
	}

	for _,x := range courses {
		if strings.Contains(x,"Go"){
			fmt.Println("Go is found in ", x,"and index is", strings.Index(x, "Go"))
		}
	}

	newString := "Go is a great programming language. Go for it!"
	fmt.Println(strings.HasPrefix(newString, "Go"))
	fmt.Println(strings.HasPrefix(newString, "Python"))
	fmt.Println(strings.HasSuffix(newString, "!"))
	fmt.Println(strings.Count(newString, "Go"))
	fmt.Println(strings.Count(newString, "Fish"))
	fmt.Println(strings.Index(newString, "Python"))
	fmt.Println(strings.LastIndex(newString, "Go"))


	if strings.Contains(newString,"Go"){
		
		// Below will Replace all as we have "-1" as the last parameter
		// Which means all

		// newString = strings.Replace(newString, "Go", "GoLang", -1)
		
		// The below would replace only one occurance
		// newString = strings.Replace(newString, "Go", "GoLang", 1)

		// Replace all occurances of Go with GoLang
		newString = strings.ReplaceAll(newString, "Go", "GoLang")
		fmt.Println(newString)

		// String Comparison
		a := "A"
		if a == "A"{
			fmt.Println("a is equal to A")
		}

		if a != "A"{
			fmt.Println("a is not equal to A")
		}

		if a == "A"{
			fmt.Println("a is equal to A")
		}

		if "A" > "B"{
			fmt.Println("A is greater than B")
		}else{
			fmt.Println("A is not greater than B")
		}

		if "Alpha" > "Absolute"{
			fmt.Println("Alpha is greater than B")
		}else{
			fmt.Println("Absolute is not greater than B")
		}

		badEmail := " me@here.com "
		badEmail = strings.TrimSpace(badEmail)
		fmt.Printf("=%s=", badEmail)
		fmt.Println()

		
	}	

	str := "alpha alpha alpha alpha alpha"
	str = replaceWith(str, "alpha", "beta", 3)
	fmt.Println(str)
	
	// Case
	myNewString := "This is a clear EXAMPLE of why we search in one case only."
	searchString := strings.ToLower(myNewString)
	if strings.Contains(searchString, "this"){
		fmt.Println("Found it!")
	}else{
		fmt.Println("Did not find it!")
	}

	// Other case functions
	fmt.Println(strings.ToLower(myNewString))
	fmt.Println(strings.ToUpper(myNewString))

	
	fmt.Println(strings.Title(strings.ToLower(myNewString)))
	

}


func replaceWith(s, old, new string, n int) string{
	// index
	i := 0

	for j := 1; j<=n; j++{
		x := strings.Index(s[i:], old)
		if x < 0{
			// we did not find it
			break
		}

		// have found it
		i = i + x
		if j == n {
			return s[:i] + new + s[i + len(old):]
		}

		i += len(old)
	}
	return s
}

