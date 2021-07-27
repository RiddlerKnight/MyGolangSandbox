package main

import "fmt"

type person struct {
	fist string
	last string
	age  int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	anAgent := secretAgent{
		person: person{
			fist: "Jame",
			last: "Bond",
			age:  32},
		ltk: true,
	}

	fmt.Println(anAgent)
	fmt.Println(
		anAgent.fist,
		anAgent.last,
		anAgent.age,
		// this is got promoted
		// this like inheritance
		fmt.Sprint(anAgent.ltk),
	)
}
