package main

import (
	"fmt"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Hello World with new linetest me now \n " +
	// 		" test again")

	// 	if i%3 == 0 {
	// 		fmt.Println("This is mod 3 : " + strconv.FormatInt(int64(i), 10))
	// 		quote.Hello()
	// 	}
	// }

	const foo = "bar"
	fmt.Println(foo)

	checkResult := 99
	fmt.Println(checkResult)

	checkResult += 55
	fmt.Println(checkResult)

	strCheckResult := "foo" // := short decoration operator
	fmt.Println(strCheckResult)

	//strCheckResult += " : " + strconv.FormatInt(int64(checkResult), 10)
	fmt.Println("==========================================")
	strCheckResult += fmt.Sprintf(" : %d", checkResult)

	fmt.Println(strCheckResult)

	var z int = 99

	fmt.Println(z)

	var a bool = true

	fmt.Println(a)

	fmt.Println("==========================================")

	justFloat := fmt.Sprintf("%.5f", 99.93333)

	fmt.Println(justFloat)

	personObj := new(person)
	n := "HH"
	personObj.name = &n
	//person.age = 50

	passMe(&personObj.name)
	fmt.Println(personObj.name)

}

type person struct {
	name *string
	age  *int
}

func passMe(str **string) {
	**str = "My NewStr"
	fmt.Println(*str)
}
