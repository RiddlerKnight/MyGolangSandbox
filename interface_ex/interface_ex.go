package main

import "fmt"

type IHumanCap interface {
	Read(value string)
	Write(value string)
}

func (s Person) Read(value string) {
	fmt.Println("Read value: ", value)
	fmt.Println("Who Read : ", s.name)
}

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
}
