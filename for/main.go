package main

import "fmt"

const youAreCool = true

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	x := 9
	for x < 10 {
		fmt.Println("X é menor que 10")
		x++
	}

	// Inicia loop infinito
	for {
		x++

		if x == 50 {
			if youAreCool {
				// Continua no loop infinito
				continue
			} else {
				break
			}
		}

		if x == 100 {
			fmt.Println("X é igual a 100")
			// Sair do loop infinito
			break
		}
	}
}
