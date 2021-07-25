package main

import "fmt"

func main() {
	m := map[string]int{
		"James": 32,
		"Penny": 27,
	}

	fmt.Println(m)
	fmt.Println("James Age: " + fmt.Sprint(m["James"]))
	fmt.Println("Penny Age: " + fmt.Sprint(m["Penny"]))
	fmt.Println("Conan Age " + fmt.Sprint(m["Conan"]))

	// Just check
	v, ok := m["Conan"]
	fmt.Println(v)
	fmt.Println(ok)
}
