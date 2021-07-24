// go run GodSpeed/Simple/TypeConvert this project

package main

import (
	"fmt"
)

var foo int
var bar string
var baz bool

func main() {
	fmt.Println("Type Convert Start")
	foo = 10
	fmt.Println(foo)

	fmt.Println("===============")
	newfoo := string(foo) // this will not work
	fmt.Println(newfoo)

	fmt.Println("===============")
	newfoo2 := fmt.Sprint(foo)
	fmt.Println(newfoo2)

	bar = "20"
	//barnew = int(bar)  // this will not work and got error
	//fmt.Println(barnew)

	newbar, newbarerror := fmt.Print(bar)
	fmt.Printf("%d", newbar)

	if newbarerror != nil {
		//error here
	}
}
