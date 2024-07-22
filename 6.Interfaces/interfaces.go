package main

import (
	"fmt"
)

// Declaración de una interfaz
type Animal interface {
	Hablar() string
}

type Perro struct{}
type Gato struct{}

// Implementación de la interfaz para Perro
func (p Perro) Hablar() string {
	return "Guau!"
}

// Implementación de la interfaz para Gato
func (g Gato) Hablar() string {
	return "Miau!"
}

func main() {
	var animal Animal

	animal = Perro{}
	fmt.Println("Perro:", animal.Hablar())

	animal = Gato{}
	fmt.Println("Gato:", animal.Hablar())

	// Interfaz vacía
	var i interface{}
	i = "Una cadena"
	fmt.Println("Interfaz vacía:", i)
	i = 42
	fmt.Println("Interfaz vacía:", i)
}
