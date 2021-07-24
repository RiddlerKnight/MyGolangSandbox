// go run GodSpeed/Simple/TypeConvert this project

package main

import (
	"fmt"
	"strconv"
)

var foo int64
var bar string
var baz bool

func main() {
	fmt.Println("Type Convert Start")
	fmt.Println("======== print int directly =======")
	foo = 10
	fmt.Println(foo)

	fmt.Println("======== convert int to string =======")
	newfoo := string(foo) // this will not work
	fmt.Println(newfoo)

	fmt.Println("======= use fmt ========")
	newfoo2 := fmt.Sprint(foo)
	fmt.Println(newfoo2)

	fmt.Println("======== use strconv =======")
	newfoo3 := strconv.FormatInt(foo, 10)
	fmt.Println(newfoo3)

	fmt.Println("======== convert string to int =======")
	bar = "20"
	//barnew = int(bar)  // this will not work and got error
	//fmt.Println(barnew)

	newbar, newbarerror := strconv.ParseInt(bar, 10, 64)
	fmt.Println(newbar)

	if newbarerror != nil {
		//error here
		panic(newbarerror)
	}

	fmt.Println("======== convert bool to string =======")
	//newbaz := string(baz) // not work
	newbaz := strconv.FormatBool(baz)
	fmt.Println(newbaz)
}
