package main

import (
	"errors"
	"fmt"
)

func main() {
	// Manejo de errores
	resultado, err := dividir(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}

	// Uso de defer, panic y recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado en el main:", r)
		}
	}()
	puedePanic()
}

func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("división por cero")
	}
	return a / b, nil
}

func puedePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado en puedePanic:", r)
		}
	}()
	panic("Esto es un panic!")
}
