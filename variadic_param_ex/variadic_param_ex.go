package main

import "fmt"

func main() {
	FooFunc(5, 7, 9, 11)
}

func FooFunc(justInt ...int) {
	for i := 0; i < len(justInt); i++ {
		fmt.Println(i)
		fmt.Println(justInt[i])
		fmt.Println("========")
	}
}
