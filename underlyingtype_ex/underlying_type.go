package main

import (
	"fmt"
	"strconv"
)

type foo string
type Bar foo
type baz []Bar
type qux baz

type Jodge struct {
	id   int64
	name string
}

func main() {

	// As String
	var justbar Bar = "My String"
	fmt.Println(justbar)
	//fmt.Print("%T", justbar)

	thejodge := Jodge{id: 50, name: "Mock"}
	fmt.Println(thejodge)
	fmt.Printf("%T \n", thejodge)

	var abaz baz = []Bar{"H1"}

	for i := int64(0); i < 10; i++ {
		//append(abaz, "H)

		//abaz = append(abaz, "Hi") //[]Bar{})
		abaz = append(abaz, Bar("H"+strconv.FormatInt(i, 10)))
		fmt.Println(abaz[i])
	}
}
