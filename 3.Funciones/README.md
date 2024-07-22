### 3. Funciones

Las funciones en Go son bloques de código reutilizables que realizan una tarea específica. Las funciones permiten organizar y estructurar el código, haciéndolo más modular y fácil de mantener. En este tema, cubriremos cómo declarar y utilizar funciones, pasar argumentos y devolver valores, funciones con múltiples valores de retorno, y funciones anónimas.

## Contenido

1. Declaración y Uso de Funciones
2. Paso de Argumentos
3. Valores de Retorno
4. Múltiples Valores de Retorno
5. Funciones Anónimas

### Declaración y Uso de Funciones

En Go, las funciones se declaran utilizando la palabra clave `func`, seguida del nombre de la función, una lista de parámetros entre paréntesis, el tipo de valor de retorno (si lo hay), y el cuerpo de la función entre llaves `{}`.

#### Ejemplo de Declaración y Uso de Funciones

```go
package main

import (
    "fmt"
)

// Declaración de una función
func saludar(nombre string) {
    fmt.Printf("Hola, %s!\n", nombre)
}

func main() {
    // Llamada a la función
    saludar("Juan")
}
```

**Explicación:**
- `func saludar(nombre string)`: Declara una función llamada `saludar` que toma un parámetro `nombre` de tipo `string`.
- `fmt.Printf("Hola, %s!\n", nombre)`: Imprime un saludo personalizado utilizando el valor del parámetro `nombre`.
- `saludar("Juan")`: Llama a la función `saludar` con el argumento `"Juan"`.

### Paso de Argumentos

Los argumentos se pasan a las funciones por valor. Esto significa que cualquier cambio hecho a los argumentos dentro de la función no afecta a los argumentos originales.

#### Ejemplo de Paso de Argumentos

```go
package main

import (
    "fmt"
)

func duplicar(valor int) {
    valor = valor * 2
    fmt.Println("Valor duplicado dentro de la función:", valor)
}

func main() {
    x := 5
    duplicar(x)
    fmt.Println("Valor original fuera de la función:", x)
}
```

**Explicación:**
- `func duplicar(valor int)`: Declara una función llamada `duplicar` que toma un parámetro `valor` de tipo `int`.
- `valor = valor * 2`: Duplica el valor del parámetro `valor` dentro de la función.
- `fmt.Println("Valor duplicado dentro de la función:", valor)`: Imprime el valor duplicado dentro de la función.
- `x := 5`: Declara una variable `x` con valor `5`.
- `duplicar(x)`: Llama a la función `duplicar` con el argumento `x`.
- `fmt.Println("Valor original fuera de la función:", x)`: Imprime el valor original de `x` fuera de la función, que no se ve afectado por los cambios dentro de la función.

### Valores de Retorno

Las funciones en Go pueden devolver valores. Se especifica el tipo de valor de retorno después de los paréntesis de la lista de parámetros.

#### Ejemplo de Valores de Retorno

```go
package main

import (
    "fmt"
)

// Declaración de una función con un valor de retorno
func suma(a int, b int) int {
    return a + b
}

func main() {
    resultado := suma(3, 4)
    fmt.Println("La suma es:", resultado)
}
```

**Explicación:**
- `func suma(a int, b int) int`: Declara una función llamada `suma` que toma dos parámetros `a` y `b` de tipo `int`, y devuelve un valor de tipo `int`.
- `return a + b`: Devuelve la suma de `a` y `b`.
- `resultado := suma(3, 4)`: Llama a la función `suma` con los argumentos `3` y `4`, y asigna el valor de retorno a la variable `resultado`.
- `fmt.Println("La suma es:", resultado)`: Imprime el resultado de la suma.

### Múltiples Valores de Retorno

Una característica poderosa de Go es que las funciones pueden devolver múltiples valores. Esto es útil para devolver un resultado junto con un valor de error, entre otros usos.

#### Ejemplo de Múltiples Valores de Retorno

```go
package main

import (
    "fmt"
    "strconv"
)

// Declaración de una función con múltiples valores de retorno
func convertirACadena(numero int) (string, error) {
    cadena := strconv.Itoa(numero)
    return cadena, nil
}

func main() {
    str, err := convertirACadena(123)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("El número convertido a cadena es:", str)
    }
}
```

**Explicación:**
- `func convertirACadena(numero int) (string, error)`: Declara una función llamada `convertirACadena` que toma un parámetro `numero` de tipo `int`, y devuelve dos valores: un `string` y un `error`.
- `cadena := strconv.Itoa(numero)`: Convierte el número a una cadena.
- `return cadena, nil`: Devuelve la cadena convertida y `nil` como valor de error.
- `str, err := convertirACadena(123)`: Llama a la función `convertirACadena` con el argumento `123`, y asigna los valores de retorno a `str` y `err`.
- `if err != nil { ... } else { ... }`: Comprueba si hay un error y maneja el resultado en consecuencia.

### Funciones Anónimas

Las funciones anónimas son funciones que no tienen un nombre y pueden ser definidas y utilizadas directamente. Son útiles para definir funciones en línea, especialmente como argumentos a otras funciones.

#### Ejemplo de Funciones Anónimas

```go
package main

import (
    "fmt"
)

func main() {
    // Declaración y uso de una función anónima
    saludo := func(nombre string) string {
        return "Hola, " + nombre + "!"
    }

    mensaje := saludo("Mundo")
    fmt.Println(mensaje)

    // Función anónima como argumento a otra función
    ejecutarFuncion := func(f func(string) string, valor string) {
        resultado := f(valor)
        fmt.Println(resultado)
    }

    ejecutarFuncion(saludo, "GoLang")
}
```

**Explicación:**
- `saludo := func(nombre string) string { ... }`: Declara una función anónima que toma un parámetro `nombre` de tipo `string` y devuelve un `string`, y la asigna a la variable `saludo`.
- `mensaje := saludo("Mundo")`: Llama a la función anónima `saludo` con el argumento `"Mundo"` y asigna el resultado a `mensaje`.
- `fmt.Println(mensaje)`: Imprime el mensaje devuelto por la función anónima.
- `ejecutarFuncion := func(f func(string) string, valor string) { ... }`: Declara una función anónima que toma una función `f` y un valor `valor` como parámetros, y la asigna a la variable `ejecutarFuncion`.
- `ejecutarFuncion(saludo, "GoLang")`: Llama a la función `ejecutarFuncion` pasando la función `saludo` y el valor `"GoLang"` como argumentos.

### Resumen

En este tema, hemos cubierto las funciones en Go, incluyendo cómo declarar y utilizar funciones, pasar argumentos y devolver valores, funciones con múltiples valores de retorno, y funciones anónimas. Las funciones son una parte fundamental de la programación en Go, y dominar su uso te permitirá escribir código más modular, reutilizable y fácil de mantener.