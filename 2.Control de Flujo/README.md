### 2. Control de Flujo

En este tema, cubrimos las estructuras de control de flujo en Go, incluyendo condicionales y bucles. Estas estructuras son esenciales para dirigir el flujo de ejecución del programa basado en condiciones y para repetir bloques de código múltiples veces.

## Contenido

1. Condicionales
2. Switch
3. Bucles

### Condicionales

Los condicionales se utilizan para ejecutar bloques de código basados en condiciones. En Go, los condicionales se implementan utilizando `if`, `else if` y `else`.

#### Ejemplo de Condicionales

```go
package main

import (
    "fmt"
)

func main() {
    x := 10

    // Condicional simple
    if x > 5 {
        fmt.Println("x es mayor que 5")
    }

    // Condicional con else
    if x < 5 {
        fmt.Println("x es menor que 5")
    } else {
        fmt.Println("x es mayor o igual a 5")
    }

    // Condicional con else if
    if x < 5 {
        fmt.Println("x es menor que 5")
    } else if x == 5 {
        fmt.Println("x es igual a 5")
    } else {
        fmt.Println("x es mayor que 5")
    }
}
```

**Explicación:**
- `if x > 5`: Comprueba si `x` es mayor que `5`.
- `if x < 5`: Comprueba si `x` es menor que `5`.
- `else if x == 5`: Comprueba si `x` es igual a `5`.
- `else`: Se ejecuta si ninguna de las condiciones anteriores es verdadera.

### Switch

El `switch` en Go es una forma más limpia y eficiente de manejar múltiples condiciones que los condicionales anidados. Se puede utilizar para simplificar la lógica de selección basada en múltiples valores posibles de una variable.

#### Ejemplo de Switch

```go
package main

import (
    "fmt"
)

func main() {
    y := 2

    switch y {
    case 1:
        fmt.Println("y es 1")
    case 2:
        fmt.Println("y es 2")
    case 3:
        fmt.Println("y es 3")
    default:
        fmt.Println("y no es 1, 2, ni 3")
    }

    // Switch sin condición (similar a if-else if-else)
    z := 4
    switch {
    case z < 3:
        fmt.Println("z es menor que 3")
    case z == 3:
        fmt.Println("z es igual a 3")
    default:
        fmt.Println("z es mayor que 3")
    }
}
```

**Explicación:**
- `switch y`: Inicia un bloque `switch` que compara `y` con varios casos (`case`).
- `case 1`: Si `y` es igual a `1`, se ejecuta este bloque.
- `default`: Se ejecuta si ninguno de los casos anteriores es verdadero.
- `switch { ... }`: Un `switch` sin condición puede reemplazar múltiples `if-else if-else` y evaluará la primera condición verdadera.

### Bucles

Los bucles permiten ejecutar un bloque de código repetidamente. En Go, el bucle más común es el `for`, que puede funcionar como un bucle `while` o `do-while` en otros lenguajes.

#### Ejemplo de Bucles

```go
package main

import (
    "fmt"
)

func main() {
    // Bucle for clásico
    for i := 0; i < 5; i++ {
        fmt.Println("Iteración", i)
    }

    // Bucle for como while
    j := 0
    for j < 5 {
        fmt.Println("Valor de j:", j)
        j++
    }

    // Bucle infinito con break
    k := 0
    for {
        fmt.Println("Bucle infinito, k:", k)
        k++
        if k == 5 {
            break
        }
    }

    // Bucle for con range
    numeros := []int{1, 2, 3, 4, 5}
    for index, value := range numeros {
        fmt.Printf("Índice: %d, Valor: %d\n", index, value)
    }
}
```

**Explicación:**
- `for i := 0; i < 5; i++`: Un bucle `for` clásico con inicialización, condición y post-incremento.
- `for j < 5`: Un bucle `for` que funciona como un bucle `while`.
- `for { ... }`: Un bucle infinito que se detiene con `break` cuando `k` es igual a `5`.
- `for index, value := range numeros`: Un bucle `for` que itera sobre un slice utilizando `range`, proporcionando el índice y el valor de cada elemento.

### Resumen

En este tema, hemos cubierto las estructuras de control de flujo en Go, incluyendo condicionales (`if`, `else if`, `else`), el uso de `switch` para manejar múltiples condiciones y diversas formas de utilizar bucles `for`. Estas herramientas permiten dirigir el flujo de ejecución de tus programas de manera eficiente y flexible. Con un buen entendimiento de estas estructuras, puedes construir programas más complejos y funcionales en Go.