package main

import (
	"fmt"
)

func main() {
	// Variables y constantes
	var variableEntera int = 10       // Declaración de una variable entera
	variableCadena := "Hola, GoLang!" // Declaración con inferencia de tipo
	const constantePi float64 = 3.14  // Declaración de una constante

	fmt.Println("Variable entera:", variableEntera)
	fmt.Println("Variable cadena:", variableCadena)
	fmt.Println("Constante Pi:", constantePi)

	// Tipos de datos
	var entero int = 42          // Tipo entero
	var flotante float64 = 3.14  // Tipo flotante
	var booleano bool = true     // Tipo booleano
	var cadena string = "GoLang" // Tipo cadena

	fmt.Println("Entero:", entero)
	fmt.Println("Flotante:", flotante)
	fmt.Println("Booleano:", booleano)
	fmt.Println("Cadena:", cadena)

	// Operadores
	suma := 10 + 20     // Operador aritmético de suma
	resta := 30 - 10    // Operador aritmético de resta
	producto := 10 * 20 // Operador aritmético de multiplicación
	division := 40 / 10 // Operador aritmético de división
	modulo := 10 % 3    // Operador aritmético de módulo

	fmt.Println("Suma:", suma)
	fmt.Println("Resta:", resta)
	fmt.Println("Producto:", producto)
	fmt.Println("División:", division)
	fmt.Println("Módulo:", modulo)

	// Operadores lógicos y comparativos
	igual := (10 == 20)       // Operador de igualdad
	diferente := (10 != 20)   // Operador de desigualdad
	mayor := (10 > 5)         // Operador mayor que
	menor := (5 < 10)         // Operador menor que
	mayorOIgual := (10 >= 10) // Operador mayor o igual que
	menorOIgual := (5 <= 10)  // Operador menor o igual que

	fmt.Println("Igual:", igual)
	fmt.Println("Diferente:", diferente)
	fmt.Println("Mayor:", mayor)
	fmt.Println("Menor:", menor)
	fmt.Println("Mayor o Igual:", mayorOIgual)
	fmt.Println("Menor o Igual:", menorOIgual)

	// Operadores lógicos
	and := (true && false) // Operador lógico AND
	or := (true || false)  // Operador lógico OR
	not := !true           // Operador lógico NOT

	fmt.Println("AND:", and)
	fmt.Println("OR:", or)
	fmt.Println("NOT:", not)
}
