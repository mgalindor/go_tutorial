package main

import (
	"fmt"
	"miModulo/maths"

	"github.com/fatih/color"
)

func main() {
	resultado := maths.Sumar(3, 4)
	fmt.Println("La suma es:", resultado)

	color.Cyan("Este es un mensaje en color cian")
}
