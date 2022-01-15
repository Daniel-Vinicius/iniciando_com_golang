package main

import "fmt"

func main() {
	a := 10

	if a > 10 {
		fmt.Println("A é maior que 10")
	}

	youAreCool := true

	if x := "Cool"; youAreCool {
		fmt.Println(x) // Cool
	}
}
