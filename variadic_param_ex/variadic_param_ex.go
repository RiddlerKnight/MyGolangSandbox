package main

import (
	"fmt"
)

func main() {
	FooFunc(5, 7, 9, 11)
}

func FooFunc(justInt ...int) {
	for i := 0; i < len(justInt); i++ {
		fmt.Println(i)
		fmt.Println(justInt[i])
		fmt.Println("========")
	}

	What(justInt...)

	sum := 0
	for i, v := range justInt {
		sum += v
		fmt.Println("Now v is " + fmt.Sprint(v) + " and i is " + fmt.Sprint(i))
	}

	fmt.Println("Total sum = ", sum)
}

func What(justAnotherInt ...int) {

}
