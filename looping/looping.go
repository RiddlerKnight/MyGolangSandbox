package main

import "fmt"

func main() {

	fmt.Println("===============")
	fmt.Println("Basic for")
	// Basic for
	for s := 0; s < 10; s++ {
		fmt.Println(s)
	}

	fmt.Println("===============")
	fmt.Println("Short for")
	/// Short for
	foo := 0
	for foo < 3 {
		fmt.Println(foo)
		foo++
	}

	fmt.Println("===============")
	fmt.Println("While for")
	/// While for
	bar := true
	i := 0
	for bar {
		if i == 10 {
			bar = false
			fmt.Println(i)
		}
		i++
	}

	fmt.Println("===============")
	fmt.Println("While true with break")
	// While true with break
	g := 0
	for {
		if g == 10 {
			fmt.Println(g)
			break
		}
		g++
	}

	fmt.Println("===============")
	fmt.Println("Basic for with continue")
	// basic for with continue
	for s := 0; s < 10; s++ {
		// Only odd number
		if s%2 == 0 {
			continue
		}
		fmt.Println(s)
	}
}
