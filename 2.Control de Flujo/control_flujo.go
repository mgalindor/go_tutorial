package main

import (
	"fmt"
)

func main() {
	// Condicionales
	x := 10
	if x > 5 {
		fmt.Println("x es mayor que 5")
	} else {
		fmt.Println("x es menor o igual a 5")
	}

	// Switch
	y := 2
	switch y {
	case 1:
		fmt.Println("y es 1")
	case 2:
		fmt.Println("y es 2")
	default:
		fmt.Println("y no es 1 ni 2")
	}

	// Bucles
	// Bucle for
	for i := 0; i < 5; i++ {
		fmt.Println("Iteración", i)
	}

	// Bucle for con range
	numeros := []int{1, 2, 3, 4, 5}
	for index, value := range numeros {
		fmt.Printf("Índice: %d, Valor: %d\n", index, value)
	}
}
