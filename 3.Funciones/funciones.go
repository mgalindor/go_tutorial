package main

import (
	"fmt"
)

// Función básica
func suma(a int, b int) int {
	return a + b
}

// Función con múltiples valores de retorno
func dividir(a, b int) (int, int) {
	cociente := a / b
	resto := a % b
	return cociente, resto
}

func main() {
	resultado := suma(3, 4)
	fmt.Println("Suma:", resultado)

	cociente, resto := dividir(10, 3)
	fmt.Println("Cociente:", cociente, "Resto:", resto)

	// Función anónima
	anonima := func(x int) int {
		return x * x
	}
	fmt.Println("Cuadrado:", anonima(5))

	// Recursividad
	fmt.Println("Factorial de 5:", factorial(5))
}

// Función recursiva
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
