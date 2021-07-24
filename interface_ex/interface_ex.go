package main

import "fmt"

type IHumanCap interface {
	Read(value string)
	Write(value string)
}

// This Look like extension
func (s Person) Read(value string) {
	fmt.Println("Read value: ", value)
	fmt.Println("Who Read : ", s.name)
}

// This Look like extension
func (s Person) Write(value string) {
	fmt.Println("Write value: ", value)
	fmt.Println("Who Write : ", s.name)
}

type Person struct {
	name string
}

func main() {
	p := Person{name: "Yuna"}
	p.Read("Read From main")
	p.Write("Write From main")

	fmt.Println("=======================")

	p2 := Person{name: "Akari"}
	p2.Read("Read From file")
	p2.Write("Write From console")

	fmt.Println("=======================")

	// There are encap with interface, they can't access name prop directly
	var p1encap IHumanCap = p
	var p2encap IHumanCap = p2

	p1encap.Read("")
	p2encap.Read("")

	p1encap.Write("")
	p2encap.Write("")
}
