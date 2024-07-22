### 6. Interfaces

En Go, las interfaces son tipos abstractos que definen un conjunto de métodos pero no implementan estos métodos. Las interfaces permiten que diferentes tipos compartan el mismo comportamiento, proporcionando una forma de lograr polimorfismo.

## Contenido

1. Definición de Interfaces
2. Implementación de Interfaces
3. Uso de Interfaces
4. Interfaces Vacías
5. Type Assertion y Type Switch

### Definición de Interfaces

Una interfaz en Go se define utilizando la palabra clave `interface`. Una interfaz contiene un conjunto de métodos que un tipo debe implementar para satisfacer la interfaz.

#### Ejemplo de Definición de Interfaces

```go
package main

import (
    "fmt"
)

// Definición de una interfaz
type Animal interface {
    Hablar() string
}
```

**Explicación:**
- `type Animal interface { ... }`: Declara una interfaz llamada `Animal` con un método `Hablar` que devuelve un `string`.

### Implementación de Interfaces

Para que un tipo satisfaga una interfaz, debe implementar todos los métodos de la interfaz. En Go, no es necesario declarar explícitamente que un tipo implementa una interfaz; si el tipo tiene los métodos requeridos, se dice que satisface la interfaz.

#### Ejemplo de Implementación de Interfaces

```go
package main

import (
    "fmt"
)

// Definición de una interfaz
type Animal interface {
    Hablar() string
}

// Definición de tipos que implementan la interfaz
type Perro struct{}
type Gato struct{}

func (p Perro) Hablar() string {
    return "Guau"
}

func (g Gato) Hablar() string {
    return "Miau"
}

func main() {
    // Creación de instancias de los tipos
    var animal Animal

    animal = Perro{}
    fmt.Println("El perro dice:", animal.Hablar())

    animal = Gato{}
    fmt.Println("El gato dice:", animal.Hablar())
}
```

**Explicación:**
- `type Perro struct{}` y `type Gato struct{}`: Declaran dos tipos, `Perro` y `Gato`.
- `func (p Perro) Hablar() string { ... }` y `func (g Gato) Hablar() string { ... }`: Implementan el método `Hablar` para los tipos `Perro` y `Gato`.
- `var animal Animal`: Declara una variable `animal` de tipo `Animal`.
- `animal = Perro{}` y `animal = Gato{}`: Asignan instancias de `Perro` y `Gato` a la variable `animal`.

### Uso de Interfaces

Las interfaces se pueden utilizar para escribir funciones genéricas que trabajan con cualquier tipo que satisfaga la interfaz.

#### Ejemplo de Uso de Interfaces

```go
package main

import (
    "fmt"
)

// Definición de una interfaz
type Describible interface {
    Describir() string
}

// Definición de un tipo que implementa la interfaz
type Persona struct {
    Nombre string
    Edad   int
}

func (p Persona) Describir() string {
    return fmt.Sprintf("%s tiene %d años", p.Nombre, p.Edad)
}

// Función que toma una interfaz como argumento
func ImprimirDescripcion(d Describible) {
    fmt.Println(d.Describir())
}

func main() {
    p := Persona{Nombre: "Juan", Edad: 30}
    ImprimirDescripcion(p)
}
```

**Explicación:**
- `type Describible interface { ... }`: Declara una interfaz `Describible` con un método `Describir`.
- `type Persona struct { ... }`: Declara una struct `Persona`.
- `func (p Persona) Describir() string { ... }`: Implementa el método `Describir` para el tipo `Persona`.
- `func ImprimirDescripcion(d Describible) { ... }`: Declara una función que toma un argumento de tipo `Describible` y llama al método `Describir`.

### Interfaces Vacías

La interfaz vacía `interface{}` puede contener cualquier tipo. Es útil cuando se necesita una función que puede aceptar argumentos de cualquier tipo.

#### Ejemplo de Interfaces Vacías

```go
package main

import (
    "fmt"
)

func imprimirTodo(v interface{}) {
    fmt.Println(v)
}

func main() {
    imprimirTodo(42)
    imprimirTodo("Hola, Go")
    imprimirTodo(true)
}
```

**Explicación:**
- `func imprimirTodo(v interface{}) { ... }`: Declara una función que toma un argumento de tipo `interface{}` y lo imprime.
- `imprimirTodo(42)`, `imprimirTodo("Hola, Go")`, `imprimirTodo(true)`: Llama a la función `imprimirTodo` con diferentes tipos de argumentos.

### Type Assertion y Type Switch

La Type Assertion se utiliza para obtener el valor concreto de una variable de tipo interfaz. El Type Switch permite realizar acciones diferentes basadas en el tipo concreto de una variable de tipo interfaz.

#### Ejemplo de Type Assertion

```go
package main

import (
    "fmt"
)

func describir(i interface{}) {
    switch v := i.(type) {
    case string:
        fmt.Printf("Soy una cadena: %s\n", v)
    case int:
        fmt.Printf("Soy un entero: %d\n", v)
    default:
        fmt.Printf("No conozco el tipo %T\n", v)
    }
}

func main() {
    describir("Hola")
    describir(42)
    describir(true)
}
```

**Explicación:**
- `switch v := i.(type) { ... }`: Utiliza un Type Switch para manejar diferentes tipos de la variable `i`.
- `case string:`, `case int:`: Maneja los casos en que `i` es de tipo `string` o `int`.
- `default`: Maneja el caso en que `i` es de un tipo desconocido.

### Resumen

En este tema, hemos cubierto las interfaces en Go, incluyendo cómo definir interfaces, implementar interfaces en tipos, usar interfaces en funciones genéricas, trabajar con interfaces vacías, y utilizar Type Assertion y Type Switch. Las interfaces son una herramienta poderosa en Go que permiten escribir código más flexible y reutilizable. Dominar el uso de interfaces te permitirá aprovechar al máximo el polimorfismo y la abstracción en tus programas en Go.