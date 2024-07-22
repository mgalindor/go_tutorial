package main

import (
    "fmt"
)

func main() {
    // Declaración de punteros
    var x int = 10
    var p *int = &x

    fmt.Println("Valor de x:", x)
    fmt.Println("Dirección de x:", p)
    fmt.Println("Valor en la dirección de p:", *p)

    // Pasar punteros a funciones
    y := 20
    fmt.Println("Antes de modificar:", y)
    modificar(&y)
    fmt.Println("Después de modificar:", y)
}

func modificar(z *int) {
    *z = 30
}
