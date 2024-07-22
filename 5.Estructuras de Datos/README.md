### 5. Estructuras de Datos

En Go, las estructuras de datos proporcionan formas eficientes de almacenar y organizar datos. En este tema, cubriremos los arrays, slices, maps y structs, que son las estructuras de datos más utilizadas en Go.

## Contenido

1. Arrays
2. Slices
3. Maps
4. Structs

### Arrays

Un array es una colección de elementos del mismo tipo con un tamaño fijo. Los elementos de un array se almacenan en ubicaciones contiguas de memoria.

#### Ejemplo de Arrays

```go
package main

import (
    "fmt"
)

func main() {
    // Declaración de un array de enteros con 5 elementos
    var numeros [5]int

    // Asignación de valores al array
    numeros[0] = 1
    numeros[1] = 2
    numeros[2] = 3
    numeros[3] = 4
    numeros[4] = 5

    // Acceso a los elementos del array
    fmt.Println("Array:", numeros)

    // Inicialización de un array con valores
    dias := [7]string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"}
    fmt.Println("Días de la semana:", dias)
}
```

**Explicación:**
- `var numeros [5]int`: Declara un array `numeros` de tipo `int` con 5 elementos.
- `numeros[i] = valor`: Asigna valores a los elementos del array.
- `dias := [7]string{"Lunes", ...}`: Declara e inicializa un array `dias` de tipo `string` con 7 elementos.

### Slices

Un slice es una estructura de datos más flexible que un array. Los slices tienen un tamaño dinámico y pueden crecer o reducirse según sea necesario. Internamente, un slice es una referencia a un array subyacente.

#### Ejemplo de Slices

```go
package main

import (
    "fmt"
)

func main() {
    // Declaración de un slice de enteros
    var numeros []int

    // Uso de la función append para agregar elementos al slice
    numeros = append(numeros, 1, 2, 3, 4, 5)
    fmt.Println("Slice:", numeros)

    // Creación de un slice utilizando un array
    diasArray := [7]string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"}
    diasSlice := diasArray[0:5] // Slice que contiene los primeros 5 días
    fmt.Println("Días laborales:", diasSlice)

    // Función len y cap
    fmt.Println("Longitud del slice:", len(diasSlice))
    fmt.Println("Capacidad del slice:", cap(diasSlice))
}
```

**Explicación:**
- `var numeros []int`: Declara un slice `numeros` de tipo `int`.
- `numeros = append(numeros, 1, 2, 3, 4, 5)`: Agrega elementos al slice `numeros` utilizando la función `append`.
- `diasSlice := diasArray[0:5]`: Crea un slice `diasSlice` que referencia los primeros 5 elementos del array `diasArray`.
- `len(diasSlice)`: Obtiene la longitud del slice.
- `cap(diasSlice)`: Obtiene la capacidad del slice.

### Maps

Un map es una colección desordenada de pares clave-valor. Los maps permiten una búsqueda rápida de valores asociados a una clave.

#### Ejemplo de Maps

```go
package main

import (
    "fmt"
)

func main() {
    // Declaración e inicialización de un map
    capitales := map[string]string{
        "España":  "Madrid",
        "Francia": "París",
        "Italia":  "Roma",
    }

    // Acceso a los valores del map
    fmt.Println("Capital de España:", capitales["España"])
    fmt.Println("Capital de Francia:", capitales["Francia"])

    // Agregar y eliminar elementos en el map
    capitales["Alemania"] = "Berlín"
    delete(capitales, "Italia")

    // Iterar sobre los elementos del map
    for pais, capital := range capitales {
        fmt.Printf("La capital de %s es %s\n", pais, capital)
    }
}
```

**Explicación:**
- `capitales := map[string]string{...}`: Declara e inicializa un map `capitales` con claves y valores de tipo `string`.
- `capitales["España"]`: Accede al valor asociado a la clave `"España"`.
- `capitales["Alemania"] = "Berlín"`: Agrega un nuevo par clave-valor al map.
- `delete(capitales, "Italia")`: Elimina el par clave-valor asociado a la clave `"Italia"` del map.
- `for pais, capital := range capitales { ... }`: Itera sobre los pares clave-valor del map.

### Structs

Una struct es una colección de campos, cada uno con su propio nombre y tipo. Las structs se utilizan para agrupar datos relacionados y pueden contener campos de diferentes tipos.

#### Ejemplo de Structs

```go
package main

import (
    "fmt"
)

// Declaración de una struct
type Persona struct {
    Nombre string
    Edad   int
    Ciudad string
}

func main() {
    // Inicialización de una struct
    juan := Persona{
        Nombre: "Juan Pérez",
        Edad:   30,
        Ciudad: "Madrid",
    }

    // Acceso a los campos de la struct
    fmt.Println("Nombre:", juan.Nombre)
    fmt.Println("Edad:", juan.Edad)
    fmt.Println("Ciudad:", juan.Ciudad)

    // Modificación de los campos de la struct
    juan.Edad = 31
    fmt.Println("Nueva edad:", juan.Edad)

    // Inicialización de una struct utilizando valores por defecto
    ana := Persona{Nombre: "Ana Gómez"}
    fmt.Println("Nombre:", ana.Nombre)
    fmt.Println("Edad:", ana.Edad)  // Valor por defecto: 0
    fmt.Println("Ciudad:", ana.Ciudad)  // Valor por defecto: ""
}
```

**Explicación:**
- `type Persona struct { ... }`: Define una struct `Persona` con tres campos: `Nombre`, `Edad` y `Ciudad`.
- `juan := Persona{ ... }`: Declara e inicializa una variable `juan` de tipo `Persona`.
- `fmt.Println("Nombre:", juan.Nombre)`: Accede al campo `Nombre` de la struct `juan`.
- `juan.Edad = 31`: Modifica el campo `Edad` de la struct `juan`.
- `ana := Persona{Nombre: "Ana Gómez"}`: Declara e inicializa una variable `ana` de tipo `Persona` utilizando valores por defecto para `Edad` y `Ciudad`.

### Resumen

En este tema, hemos cubierto las estructuras de datos fundamentales en Go, incluyendo arrays, slices, maps y structs. Los arrays y slices son útiles para almacenar colecciones de elementos del mismo tipo, mientras que los maps permiten asociar claves y valores para una búsqueda rápida. Las structs proporcionan una forma flexible de agrupar datos relacionados. Comprender y utilizar estas estructuras de datos te permitirá escribir programas más eficientes y organizados en Go.