package main

import "fmt"

func main() {
	name := "Daniel"

	switch name {
	case "Wesley":
		{
			fmt.Println("Hey Wesley")
		}

	case "Daniel":
		{
			fmt.Println("Hey Daniel")
		}

	default:
		{
			fmt.Println("Hey bro!")
		}
	}
}
