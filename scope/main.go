package main

import "fmt"

var y = 20

func main() {
	x := 10
	fmt.Println(x)
	printZ()
	printY()
}

func printY() {
	fmt.Println(y)
}
