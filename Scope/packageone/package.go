package packageone

// As "privateVar" has a smaller "p" which means this variable is available in
// the package only
var privateVar = "I am private"

// As "PublicVar" has a Capital "P" which means this variable is getting exported
// in the Go world
var PublicVar = "I am public (or Exported)"

// As "notExported" has a smaller "p" which means this variable is available in
// the package only
func notExported(){

}

// As "Exported" has a Capital "P" which means this variable is getting exported
// in the Go world
func Exported(){

}
