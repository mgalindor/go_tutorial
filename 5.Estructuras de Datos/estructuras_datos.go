package main

import (
	"fmt"
)

func main() {
	// Arrays
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)

	// Slices
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)
	slice = append(slice, 6)
	fmt.Println("Slice después de append:", slice)

	// Maps
	dict := map[string]int{"uno": 1, "dos": 2, "tres": 3}
	fmt.Println("Map:", dict)
	fmt.Println("Valor de 'dos':", dict["dos"])

	persona := Persona{Nombre: "Juan", Edad: 25}
	fmt.Println("Struct:", persona)

	// Métodos en structs
	persona.saludar()
}

// Structs
type Persona struct {
	Nombre string
	Edad   int
}

func (p Persona) saludar() {
	fmt.Println("Hola, mi nombre es", p.Nombre)
}
