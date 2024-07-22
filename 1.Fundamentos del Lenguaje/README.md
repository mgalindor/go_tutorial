### 1. Fundamentos del Lenguaje

En este tema, cubrimos los conceptos básicos del lenguaje Go, incluyendo la sintaxis, declaración de variables y constantes, tipos de datos y operadores.

## Contenido

1. Introducción a Go
2. Declaración de Variables y Constantes
3. Tipos de Datos
4. Operadores

### Introducción a Go

Go es un lenguaje de programación de código abierto desarrollado por Google. Está diseñado para ser simple, eficiente y adecuado para el desarrollo de software a gran escala. Go combina la eficiencia y seguridad de un lenguaje compilado como C con la velocidad de desarrollo de un lenguaje dinámico como Python.

### Declaración de Variables y Constantes

En Go, las variables pueden ser declaradas utilizando la palabra clave `var` o el operador de declaración corta `:=`. Las constantes se declaran utilizando la palabra clave `const`.

#### Ejemplo de Declaración de Variables y Constantes

```go
package main

import (
    "fmt"
)

func main() {
    // Declaración de variables
    var variableEntera int = 10         // Declaración explícita con tipo
    variableCadena := "Hola, GoLang!"   // Declaración con inferencia de tipo

    // Declaración de constantes
    const constantePi float64 = 3.14    // Declaración de una constante

    // Imprimir variables y constantes
    fmt.Println("Variable entera:", variableEntera)
    fmt.Println("Variable cadena:", variableCadena)
    fmt.Println("Constante Pi:", constantePi)
}
```

**Explicación:**
- `var variableEntera int = 10`: Declara una variable `variableEntera` de tipo `int` con valor `10`.
- `variableCadena := "Hola, GoLang!"`: Declara una variable `variableCadena` con inferencia de tipo `string`.
- `const constantePi float64 = 3.14`: Declara una constante `constantePi` de tipo `float64` con valor `3.14`.

### Tipos de Datos

Go tiene varios tipos de datos básicos, incluyendo enteros, flotantes, booleanos y cadenas. 

#### Ejemplo de Tipos de Datos

```go
package main

import (
    "fmt"
)

func main() {
    // Tipos de datos básicos
    var entero int = 42              // Tipo entero
    var flotante float64 = 3.14      // Tipo flotante
    var booleano bool = true         // Tipo booleano
    var cadena string = "GoLang"     // Tipo cadena

    // Imprimir tipos de datos
    fmt.Println("Entero:", entero)
    fmt.Println("Flotante:", flotante)
    fmt.Println("Booleano:", booleano)
    fmt.Println("Cadena:", cadena)
}
```

**Explicación:**
- `var entero int = 42`: Declara una variable `entero` de tipo `int`.
- `var flotante float64 = 3.14`: Declara una variable `flotante` de tipo `float64`.
- `var booleano bool = true`: Declara una variable `booleano` de tipo `bool`.
- `var cadena string = "GoLang"`: Declara una variable `cadena` de tipo `string`.

### Operadores

Los operadores en Go incluyen operadores aritméticos, operadores de comparación y operadores lógicos.

#### Operadores Aritméticos

Los operadores aritméticos se utilizan para realizar operaciones matemáticas básicas como suma, resta, multiplicación y división.

```go
package main

import (
    "fmt"
)

func main() {
    // Operadores aritméticos
    suma := 10 + 20                // Operador de suma
    resta := 30 - 10               // Operador de resta
    producto := 10 * 20            // Operador de multiplicación
    division := 40 / 10            // Operador de división
    modulo := 10 % 3               // Operador de módulo

    // Imprimir resultados
    fmt.Println("Suma:", suma)
    fmt.Println("Resta:", resta)
    fmt.Println("Producto:", producto)
    fmt.Println("División:", division)
    fmt.Println("Módulo:", modulo)
}
```

**Explicación:**
- `suma := 10 + 20`: Realiza la suma de `10` y `20`.
- `resta := 30 - 10`: Realiza la resta de `30` y `10`.
- `producto := 10 * 20`: Realiza la multiplicación de `10` y `20`.
- `division := 40 / 10`: Realiza la división de `40` entre `10`.
- `modulo := 10 % 3`: Calcula el módulo de `10` dividido por `3`.

#### Operadores de Comparación

Los operadores de comparación se utilizan para comparar dos valores. Los resultados de estas operaciones son booleanos (`true` o `false`).

```go
package main

import (
    "fmt"
)

func main() {
    // Operadores de comparación
    igual := (10 == 20)             // Operador de igualdad
    diferente := (10 != 20)         // Operador de desigualdad
    mayor := (10 > 5)               // Operador mayor que
    menor := (5 < 10)               // Operador menor que
    mayorOIgual := (10 >= 10)       // Operador mayor o igual que
    menorOIgual := (5 <= 10)        // Operador menor o igual que

    // Imprimir resultados
    fmt.Println("Igual:", igual)
    fmt.Println("Diferente:", diferente)
    fmt.Println("Mayor:", mayor)
    fmt.Println("Menor:", menor)
    fmt.Println("Mayor o Igual:", mayorOIgual)
    fmt.Println("Menor o Igual:", menorOIgual)
}
```

**Explicación:**
- `igual := (10 == 20)`: Comprueba si `10` es igual a `20`.
- `diferente := (10 != 20)`: Comprueba si `10` es diferente de `20`.
- `mayor := (10 > 5)`: Comprueba si `10` es mayor que `5`.
- `menor := (5 < 10)`: Comprueba si `5` es menor que `10`.
- `mayorOIgual := (10 >= 10)`: Comprueba si `10` es mayor o igual a `10`.
- `menorOIgual := (5 <= 10)`: Comprueba si `5` es menor o igual a `10`.

#### Operadores Lógicos

Los operadores lógicos se utilizan para combinar múltiples condiciones lógicas.

```go
package main

import (
    "fmt"
)

func main() {
    // Operadores lógicos
    and := (true && false)          // Operador lógico AND
    or := (true || false)           // Operador lógico OR
    not := !true                    // Operador lógico NOT

    // Imprimir resultados
    fmt.Println("AND:", and)
    fmt.Println("OR:", or)
    fmt.Println("NOT:", not)
}
```

**Explicación:**
- `and := (true && false)`: Realiza una operación lógica AND entre `true` y `false`.
- `or := (true || false)`: Realiza una operación lógica OR entre `true` y `false`.
- `not := !true`: Realiza una operación lógica NOT sobre `true`.

### Resumen

En este tema, hemos cubierto los conceptos básicos del lenguaje Go, incluyendo cómo declarar variables y constantes, los tipos de datos básicos y los operadores aritméticos, de comparación y lógicos. Estos fundamentos son esenciales para cualquier desarrollo en Go y proporcionan una base sólida para explorar temas más avanzados en el lenguaje.