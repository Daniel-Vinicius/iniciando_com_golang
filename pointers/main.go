package main

import "fmt"

func addOneToNumber(a int) int {
	a = a + 1
	return a
}

// addOneToNumberWithPointers Receive by parameter the reference in memory like this: &two
func addOneToNumberWithPointers(a *int) int {
	*a = *a + 1
	return *a
}

func main() {
	eight := 8
	fmt.Println(addOneToNumber(eight))
	fmt.Println(eight)

	two := 2
	fmt.Println(addOneToNumberWithPointers(&two))
	fmt.Println(two)

	x := 10
	fmt.Println(&x) // 0xc00012a000

	y := &x
	fmt.Println(y)  // 0xc00012a000
	fmt.Println(*y) // 10

	*y = 20
	fmt.Println(x) // 20

	var z *int = &x
	fmt.Println(z)  // 0xc00012a000
	fmt.Println(*z) // 20
}
