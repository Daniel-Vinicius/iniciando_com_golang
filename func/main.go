package main

import (
	"fmt"
	"strings"
)

func multiplicate(number int, times int) int {
	return number * times
}

func namedReturnFunction(a string) (x string) {
	x = a
	return
}

func moreReturns(a string, b string) (string, string) {
	return a, b
}

func variadicFunc(x ...int) int {
	var res int

	for _, v := range x {
		res += v
	}

	return res
}

func main() {
	fmt.Println(multiplicate(1430, 3))
	fmt.Println(namedReturnFunction("Daniel"))

	firstName, lastName := moreReturns("Daniel", "Vinícius")
	fmt.Println(firstName, lastName)

	total := variadicFunc(3, 5, 8, 20)
	fmt.Println(total)

	// Anonymous function
	myName := "Daniel Vinícius"

	extractLastName := func(name string) string {
		lastName := strings.Split(name, " ")[1]
		return lastName
	}

	fmt.Println(extractLastName(myName))
}
